package main

import (
	"BLCpublicchain_go/blc"
	"fmt"
	"BLCpublicchain_go/flag"
	"time"
	"BLCpublicchain_go/boltdb"
)

func main()  {

	//part2_flagCommand()
	part1_blockchain()


}

func part2_flagCommand()  {
	flag.GetCommandInputArgs()
	fmt.Println("")
	flag.SetFlag()
}

/**
	第一部分：
	如何创建/新增创世区块，
	如何创建区块链，以及将区块添加到区块链上
	如何计算hash
 */
func part1_blockchain()  {
	flagBlcParams := flag.GetFlagForBlockChain()

	//创建创世区块
	blockchain := blc.CreateGenesisBlockChain()
	//新增区块
	if flagBlcParams.AddTransitionData != "" {
		blockchain.AddBlock([]byte(flagBlcParams.AddTransitionData))
	}
	//blockchain.AddBlock([]byte("transfer A to B 100 bitcoin"))
	//blockchain.AddBlock([]byte("transfer B to C 50 bitcoin"))
	//blockchain.AddBlock([]byte("transfer B to D 35 bitcoin"))
	//blockchain.AddBlock([]byte("transfer B to B 15 bitcoin"))

	//if flagBlcParams.Printchain {
		//显示所有区块信息
		for _, block := range blockchain.GetBlockChain(){
			fmt.Println("Data ",string(block.Data))
			fmt.Printf("prev hash %x \n", block.PrevHash)
			fmt.Printf("curr hash %x \n",block.Hash)

			intTimestamp := block.Timestampe
			//时间转化为string，layout必须为 "2006-01-02 15:04:05"
			strTime := time.Unix(intTimestamp, 0).Format("2006-01-02 15:04:05")

			fmt.Printf("create time %s \n", strTime)
			fmt.Printf("nonce %v \n ", block.Nonce)

			fmt.Println()
		}


		fmt.Println("------序列化--------")
		//test 序列化
		block := blockchain.GetBlockChain()[1]
		serializeBlockBytes := block.EnSerialize()
		fmt.Println(serializeBlockBytes)


		fmt.Println("------boltdb--------")
		//create bolt db
		dbManger := boltdb.OpenDB("bc.db")

		defer dbManger.CloseDB()
		//insert data to db
		dbManger.InsertData("blockchain", []byte(block.Hash[:]), serializeBlockBytes)

		//select data from db
		selectDataFromDb := dbManger.SelectData("blockchain", []byte(block.Hash[:]))


		fmt.Println("------反序列化--------")
		//反序列化
		block = block.DeSerialize(selectDataFromDb)
		fmt.Printf("hash: %x \n",block.Hash)
		fmt.Printf("hash: %s \n",block.Data)
		fmt.Println()
	//}



}