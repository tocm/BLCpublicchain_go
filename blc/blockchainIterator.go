package blc

import (
	"BLCpublicchain_go/boltdb"
	"BLCpublicchain_go/utils"
)

/**
	区块链区块迭代器结构
 */
type BlockchainIterator struct {
	CurrentHash [32]byte //当前遍历的hash
	DBManager  *boltdb.DBManger //对应数据库
}


/*
	创建迭代器
 */
func CreateIterator(bchain *Blockchain) *BlockchainIterator {
	//创建迭代器对象，先把当前最后区块hash赋值给迭代器中的当前hash，然后往上追朔全部区块。
	return &BlockchainIterator{bchain.lastBlockHash,bchain.dbManager}
}

/**
	遍历区块
 */
func (bchainIterator *BlockchainIterator)Next() *Block {
	//判断当前区块是否已经是创世区块prevHash=0，如果是直接返回nil
	if utils.IsCompGenerateBlock(bchainIterator.CurrentHash[:]) {
		//结束遍历
		return nil
	}
	//通过最后那块hash 去查找该块中的上一块hash
	blockBytes := bchainIterator.DBManager.SelectData(boltdb.DB_TABLE_NAME_BLOCKS, bchainIterator.CurrentHash[:])
	//反序列化
	curBlock := DeSerialize(blockBytes)
	if curBlock != nil {
		//更新遍历器结构变量hash值，移动上一区块的prev hash
		bchainIterator.CurrentHash = curBlock.PrevHash
	}
	return curBlock
}


