package blc

import (
	"BLCpublicchain_go/boltdb"
	"fmt"
	"BLCpublicchain_go/utils"
)

/**
	定义区块链结构
	需要将结构信息保存到数据库中，避免只保存在内存，下次启动数据丢失
 */
type Blockchain struct {
	//blocks [] *Block //存放所有区块数组链
	lastBlockHash [32]byte //记录最后那个hash， 为了可以找到链中其它的hash
	dbManager *boltdb.DBManger //存放block database

}

/**
	初始化blockchain
 */
func Init() *Blockchain  {
	//创建区块链对象
	blockchain := new(Blockchain)
	//创建并打开数据库
	blockchain.dbManager = boltdb.OpenDB(boltdb.DB_NAME)

	//暂时没有任何块数据
	if blockchain.dbManager.IsExistDBTable(boltdb.DB_TABLE_NAME_TIP) == false {
		return blockchain
	}

	//如果已经有区块数据，将当前区块hash放到内存
	if blockchain.dbManager != nil{
		//将数据库中值取出放到内存
		 lasthash:= blockchain.dbManager.SelectData(boltdb.DB_TABLE_NAME_TIP,[]byte(boltdb.DB_TABLE_NAME_TIP_KEY_LASTBLOCKHASH))
		 copy(blockchain.lastBlockHash[:32], lasthash)

	}

	return blockchain
}

/**
	创建创世区块链
 */
func (bchain *Blockchain)CreateGenesisBlockChain()  {

	if bchain.dbManager.IsExistDBTable(boltdb.DB_TABLE_NAME_TIP) {
		fmt.Println("The Genesis block has been existed!")
		return
	}

	fmt.Println("todo Create Genesis Block ")
	//创建创世区块
	block := CreateGenesisBlock();

	//序列化block结构
	blockBytes := block.EnSerialize()
	//取出hash 当作key
	bchain.lastBlockHash = block.Hash

	//插入block数据到db
	bchain.dbManager.InsertData(boltdb.DB_TABLE_NAME_BLOCKS, bchain.lastBlockHash[:], blockBytes);
	//保存最后hash到db
	bchain.dbManager.InsertData(boltdb.DB_TABLE_NAME_TIP, []byte(boltdb.DB_TABLE_NAME_TIP_KEY_LASTBLOCKHASH), bchain.lastBlockHash[:])

	//blockchain.blocks = append(blockchain.blocks, block)
}


/**
	新增区块到区块链上
 */
func (bchain *Blockchain) AddBlock(data []byte, txs[] *Transaction)  {

	if txs == nil {
		return;
	}

	//获取链中最尾的块
	lastBlockHash := bchain.lastBlockHash
	//取出上个区块的hash
	prevHash := lastBlockHash
	//取出上一个区块
	lastBlock := bchain.GetBlockFromDB(prevHash)
	//取出上一区块的index
	blockIndex := lastBlock.Index

	//创建新区块
	newBlock := NewBlock(blockIndex+1, data, prevHash, txs)
	//将最新区块的hash 放到内存区块结构变量
	bchain.lastBlockHash = newBlock.Hash
	//准备放到数据库：先序列化
	serializeBlock := newBlock.EnSerialize()
	bchain.dbManager.InsertData(boltdb.DB_TABLE_NAME_BLOCKS, newBlock.Hash[:], serializeBlock)
	bchain.dbManager.InsertData(boltdb.DB_TABLE_NAME_TIP, []byte(boltdb.DB_TABLE_NAME_TIP_KEY_LASTBLOCKHASH), bchain.lastBlockHash[:])

}

func (bchain *Blockchain) GetBlockchainDB() (hash[32]byte, db *boltdb.DBManger) {
	return bchain.lastBlockHash, bchain.dbManager
}


/*
	从数据库中取出指定区块
	@params hash []byte  区块hash
 */
func (bchain *Blockchain)GetBlockFromDB(hash [32]byte) *Block {
	if bchain.dbManager != nil {
		//先从数据库中取出
		blockBytes := bchain.dbManager.SelectData(boltdb.DB_TABLE_NAME_BLOCKS, hash[:])
		//反序列化
		block := DeSerialize(blockBytes)
		return block
	}
	return nil
}

func (bchain *Blockchain)GetBalance(walletAddress string) int64  {

	var availableTxs[] *Transaction;//存放链上所有涉及当前用户钱包地址的交易数据

	var spentTxInputs[] *TXInput; //存放当前用户已经存在的交易输入

	var balance int64; //存放余额


	//1. 遍历所有区块并找出相关的交易数据存放在新的一个数组上
	if bchain != nil {
		bcIterator := CreateIterator(bchain)
		for {
			block := bcIterator.Next()

			if block == nil {
				break
			}

			//取出当前区块中的所有交易
			transactions := block.Transactions

			//查找与当前钱包地址相关的所有区块交易
			for _, tx := range transactions {
				//判断当前地址是否可用
				isEnableTxs := tx.IsEnableTransaction(walletAddress)
				if isEnableTxs {
					//fmt.Println("available: ", tx.TxHash)
					//重新放到一个交易
					availableTxs = append(availableTxs, tx);

					//遍历TxInput当前交易中已花费的交易输入
					for _,spentTx := range tx.TxIns {
						if spentTx.TxHash != nil {
							//fmt.Println("spent tx inputs: ", spentTx.TxHash)
							//把已花费的交易输入数据存到新数组
							spentTxInputs = append(spentTxInputs, spentTx)
						}
					}

				}
			}
		}
	}

	//2. 处理与用户地址相关的可用交易块，找出未花费的交易块
	//Find_UnSpentTxs
	for _, availabTrx := range availableTxs {
		var isSpentTxs bool
		trxhash := availabTrx.TxHash
		//fmt.Printf("getBalance s3 trx hash %x \n ", trxhash)
		//从已花费交易块中数组中的TxInput中查找是否有对应的输入，如果有即已被花费，否则即为未花费
		for _, trx := range spentTxInputs {
			//fmt.Printf("getBalance s4 compare trxhash %x \n ", trx.TxHash)
			if utils.CompareHash(trxhash, trx.TxHash) {
				//确认当前块是已花费的input txHash
				isSpentTxs = true
			}
		}

		//如果为已花费交易块，即跳出
		if isSpentTxs {
			fmt.Println("continue ")
			continue
		}

		//找出用户对应output的余额
		trxOutputs := availabTrx.TxOuts
		for _,output := range trxOutputs {
			//fmt.Println("count the user : ", output.WalletAddress)
			if output.WalletAddress == walletAddress {
				//如果满足就累计余额
				balance += output.Amount
				//fmt.Printf("User: %s, balance: %d",walletAddress, balance)
			}
		}

	}

	return balance


}

