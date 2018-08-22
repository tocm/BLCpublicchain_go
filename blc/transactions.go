package blc

import (
	"bytes"
	"encoding/gob"
	"log"
	"crypto/sha256"
	"fmt"
	"BLCpublicchain_go/utils"
)

/**
	交易结构
 */
type Transaction struct {
	TxHash []byte //交易hash
	TxIns[] *TXInput //交易输入，可包含多笔输入
	TxOuts[] *TXOutput //交易输出，可包含多笔输出
}

/**
	创建交易hash
 */
func (trax *Transaction)CreateTxHash() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(trax)
	if err != nil {
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())
	return hash[:]
}

/**
	创建创世区块交易
 */
func (trax *Transaction)CreateCoinbaseTransaction(walletAddress string)  {

	//交易输入【创始区块来自的输入是无】
	txinput := &TXInput{nil, -1, walletAddress}
	//交易输出，创世总的发行币数
	txout := &TXOutput{100000, walletAddress}
	trax.TxHash = trax.CreateTxHash()

	fmt.Printf("new transaction %s \n", walletAddress)
	trax.TxIns = append(trax.TxIns, txinput)
	trax.TxOuts = append(trax.TxOuts, txout)

}

/**
	判断该交易块是否与传入钱包地址有关易关系
 */
func (trax *Transaction)IsEnableTransaction(walletAddress string) bool {
	//fmt.Printf("IsEnableTransaction : %x \n",walletAddress)
	//先判断txInputs 是否有相关的交易
	for _,input := range trax.TxIns{
		if input.WalletAddress == walletAddress {
			//fmt.Println("IsEnableTransaction txInputs return true")
			return true
		}
	}
	//fmt.Println("-----TxOuts Begin----")
	//如果txInputs没有相关交易，再找outpus
	for _,output := range trax.TxOuts {
		if utils.CompareHash([]byte(output.WalletAddress), []byte(walletAddress)) {
			//fmt.Println("IsEnableTransaction txOuts return true")
			return true
		}
	}
	return false
}


/**
	获得未花费的交易
 */
func (trax *Transaction)GetSpentTxs(curTxHash string)   {

}


/**
	找出所有tx input交易输入并判断其是否与传入钱包地址有相关
 */
func (trax *Transaction)IsEnableTransactionIns(walletAddress string) (bool, []*TXInput) {
	var isEnableTxIns = false
	var txInsArrays[] *TXInput
	for _,txins := range trax.TxIns{
		if txins.WalletAddress == walletAddress {
			isEnableTxIns = true
			txInsArrays = append(txInsArrays, txins)
		}
	}

	return isEnableTxIns, txInsArrays
}

/**
	找出所有tx outs交易输出并判断其是否与传入钱包地址有相关
 */
func (trax *Transaction)IsEnableTransactionOuts(walletAddress string) (bool, []*TXOutput){
	var isEableTxOuts = false
	var txOutsArrays[] *TXOutput
	for _,txouts := range trax.TxOuts{
		if txouts.WalletAddress == walletAddress {
			isEableTxOuts = true
			txOutsArrays = append(txOutsArrays, txouts)
		}
	}

	return isEableTxOuts, txOutsArrays
}

/**
	创建创世区块交易
	初始化币的数量
 */
func CreateGenesisTransactions(walletAddress string) [] *Transaction {
	fmt.Println("CreateGenesisTransactions", walletAddress)
	var txs[] *Transaction
	trax := new(Transaction)
	trax.CreateCoinbaseTransaction(walletAddress)
	//把创世交易添加到交易数组
	txs = append(txs, trax)
	return txs
}