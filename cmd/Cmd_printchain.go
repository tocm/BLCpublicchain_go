package cmd

import (
	"fmt"
	"BLCpublicchain_go/blc"
)

func (cmdParams *CmdParams)PrintChain()  {

	if cmdParams.Printchain {
		fmt.Println("------所有区块信息--------")
		bchainIterator := blc.CreateIterator(cmdParams.BlockChain)
		for {
			block := bchainIterator.Next()
			if block == nil {
				break
			}
			//打印
			block.ShowBlockInfo()
		}
		fmt.Println()
	}
}
