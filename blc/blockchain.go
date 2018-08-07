package blc

import (
	"BLCpublicchain_go/boltdb"
	"fmt"
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
func (bchain *Blockchain) AddBlock(data []byte)  {

	/*
	旧的方式
	//get prevHash
	chainLen := len(bchain.blocks)
	//获取链中最尾的块
	lastBlock := bchain.blocks[chainLen - 1]

	//取出上个区块的hash
	prevHash := lastBlock.Hash
	创建新区块
	newBlock := NewBlock(data, prevHash)
	新区块追加到链尾
	bchain.blocks = append(bchain.blocks, newBlock)
	*/

	//获取链中最尾的块
	lastBlockHash := bchain.lastBlockHash
	//取出上个区块的hash
	prevHash := lastBlockHash

	//创建新区块
	newBlock := NewBlock(data, prevHash)
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




