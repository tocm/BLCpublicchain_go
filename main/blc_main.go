package main

import "BLCpublicchain_go/blc"

func main()  {

	part1_blockchain()

}

/**
	第一部分：
	如何创建/新增创世区块，
	如何创建区块链，以及将区块添加到区块链上
	如何计算hash
 */
func part1_blockchain()  {
	//创建创世区块
	blockchain := blc.CreateGenesisBlockChain()
	//新增区块
	blockchain.AddBlock([]byte("transfer A to B 100 bitcoin"))
	blockchain.AddBlock([]byte("transfer B to C 50 bitcoin"))
	//blockchain.AddBlock([]byte("transfer B to D 35 bitcoin"))
	//blockchain.AddBlock([]byte("transfer B to B 15 bitcoin"))
	//显示所有区块信息
	blockchain.ShowAllBlock()

}