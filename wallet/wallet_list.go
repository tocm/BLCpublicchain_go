package wallet

import (
	"fmt"
	"BLCpublicchain_go/storge"
	"bytes"
	"encoding/gob"
	"log"
	"crypto/elliptic"
)

//所有钱包结构
type WalletList struct {
	WalletMap map[string] *Wallet
}

/**
	创建钱包集合
 */
func GetWalletMaps() *WalletList {
	var walletList *WalletList
	if storge.IsExistFile(storge.WALLETS_ADDRESS_FILE) == false {
		walletList = new(WalletList)
		//初始化Map
		walletList.WalletMap = make(map[string] *Wallet)
		return walletList
	}

	walletListBytes := storge.ReadFile(storge.WALLETS_ADDRESS_FILE)
	//进行反序列化
	walletList = deSerialize(walletListBytes)
	return walletList
}

/**
	创建新钱包地址
 */
func (walletList *WalletList)CreateWallet() []byte {

	//创建一个钱包地址hash
	wallet := NewWallet()
	//取出地址作为map key使用
	walletAddress := wallet.GetWalletAddress()
	//创建新钱包地址并添加到集合中
	strWalletAddress := string(walletAddress)
	//新钱包存在map中
	walletList.WalletMap[strWalletAddress] = wallet

//	fmt.Println("createWallet map : ",walletList.WalletMap)

	walletsBytes := enSerialize(walletList)
	storge.SaveFile(storge.WALLETS_ADDRESS_FILE, walletsBytes)

	return walletAddress
}

/**
 	查找钱包
 */
func (walletList *WalletList)GetWallet(walletAddress string) *Wallet {
	wallet := walletList.WalletMap[walletAddress]
	return wallet
}

func (walletList *WalletList) ListWallets()  {
	fmt.Println("All Wallet address: ")
	for key,_ := range walletList.WalletMap {
		fmt.Println(key)
	}

}


// 将结构序列化成字节数组
func enSerialize(wallets *WalletList) []byte {
	var result bytes.Buffer

	//注册序列化接口
	// 注册的目的，是为了，可以序列化任何类型
	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(wallets)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化
func deSerialize(walletsBytes []byte) *WalletList {
	var wallets WalletList
	// 注册的目的，是为了，可以序列化任何类型
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(walletsBytes))
	err := decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}
	return &wallets
}