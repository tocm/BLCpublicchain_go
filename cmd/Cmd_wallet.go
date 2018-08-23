package cmd

import (
	"BLCpublicchain_go/wallet"
	"fmt"
)

func (cmd *CmdParams) CreateNewWallet()  {

	wallets := wallet.GetWalletMaps()
	walletHash := wallets.CreateWallet()

	fmt.Println("new wallet address ", string(walletHash))

}

func (cmd *CmdParams) ListWalletAddress()  {
	wallets := wallet.GetWalletMaps()
	wallets.ListWallets()

}