package cmd

import (
	"BLCpublicchain_go/wallet"
	"fmt"
)

/**
	取得余额
 */
func (cmd *CmdParams)GetBalance()  {
	if cmd.GetBalanceAddress != "" {

		isCorrectAddress := wallet.VerifyWalletAddress([]byte(cmd.GetBalanceAddress))
		if isCorrectAddress {
			balance := cmd.BlockChain.GetBalance(cmd.GetBalanceAddress)
			fmt.Printf("Wallet %s ---> balance: %d \n",cmd.GetBalanceAddress, balance)
		}
	}
}