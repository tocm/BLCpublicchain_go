package wallet
//
//import (
//	"bytes"
//	"encoding/gob"
//	"log"
//	"crypto/sha256"
//)
//
//type Wallet struct {
//	email string
//	nickName string
//	password string
//	Hash []byte
//}
//
//func CreateWallet(mail string, nickName string, pwd string) *Wallet {
//	wallet := new(Wallet)
//	wallet.email = mail;
//	wallet.nickName = nickName
//	wallet.password = pwd
//	wallet.Hash = wallet.createWalletHash()
//	return wallet
//}
//
//
///**
//	创建交易hash
// */
//func (w *Wallet)createWalletHash() []byte {
//	var result bytes.Buffer
//	encoder := gob.NewEncoder(&result)
//	err := encoder.Encode(w)
//	if err != nil {
//		log.Panic(err)
//	}
//	hash := sha256.Sum256(result.Bytes())
//	return hash[:]
//}