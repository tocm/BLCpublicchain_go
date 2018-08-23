package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"crypto/sha256"
	 "golang.org/x/crypto/ripemd160"
	"fmt"
	"BLCpublicchain_go/crypto"

	"bytes"
)

const VERSION_NUM  = byte(0x00) //1个字节
const CHECKSUM_BIT_LEN  = 4 //4个字节长度
const WALLET_ADDRESS_LEN  = 25 // VERSION(1字节) + ripemd160(20字节) + checksum(4字节)

/**
	钱包结构，一个私钥（椭圆曲线算法生成）和公钥
 */
type Wallet struct {

	//椭圆曲线算法生成私有key
	PrivateKey ecdsa.PrivateKey
	//公有key 由私有key生成
	PublicKey []byte

}

/*
	s1: 创建钱包
 */
func NewWallet() *Wallet {
	//s1 生成private key
	privKey, pubKey := newPairKey()
	return &Wallet{privKey, pubKey}
}


/**
	步骤1
	椭圆曲线算法生成私有钥,和公有钥
 */
func newPairKey() (ecdsa.PrivateKey, []byte) {
	//椭圆曲线算法生成私有钥
	curveAlgorithms := elliptic.P256()
	privKey, err := ecdsa.GenerateKey(curveAlgorithms,rand.Reader)

	if err != nil {
		log.Panic(err)
	}
	//私钥再生成公钥
	publicKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)

	return *privKey,publicKey
}

/**
	得到一个真实的比特币地址需要以下步骤：
	1.取得公钥PubKey
	2.使用 RIPEMD160(SHA256(PubKey)) 哈希算法，取公钥并对其哈希两次
	3.给哈希加上地址生成算法版本的前缀
	4.对于第二步生成的结果，使用 SHA256(SHA256(payload)) 再哈希，计算校验和。校验和是结果哈希的前四个字节。
	5.将校验和附加到 version+PubKeyHash 的组合中。
	6.使用 Base58 对 version+PubKeyHash+checksum 组合进行编码。

 */
func (w *Wallet) GetWalletAddress() []byte {
	//1: 取公钥
	pubKey := w.PublicKey
	//2: 两次hash算法RIPEMD160(SHA256(PubKey)) 得到 20字节长度
	ripemd160Hash := hashRipemd160_sha256(pubKey)

	//3: hash160 + version = 21字节
	ripemd160_versionHash := append([]byte{VERSION_NUM},ripemd160Hash...)

	//4: 两次hash payload 得到4字节长度
	checksumHashBytes := hashCheckSum(ripemd160_versionHash)

	//5. 组合 version+PubKeyHash+checksum = 25个字节
	addressBytes := append(ripemd160_versionHash,checksumHashBytes...)

	//6 进行Base58算法编码，编码后得到的hash长度是 34字符（17字节，136位）
	wallet_address := crypto.Base58Encode(addressBytes)
	fmt.Println("create wallet address len = ",len(wallet_address))
	fmt.Printf("create wallet address 0x %x \n",wallet_address)
	fmt.Printf("create wallet address string %s \n",wallet_address)
	return wallet_address
}

/**
	2: 两次hash算法RIPEMD160(SHA256(PubKey))
 */
func hashRipemd160_sha256(pubKey []byte) []byte {
	//sha256
	hash256 := sha256.New()
	hash256.Write(pubKey)
	hash256Data := hash256.Sum(nil)

	//ripemd160
	hash160 := ripemd160.New()
	hash160.Write(hash256Data)
	hash160Data := hash160.Sum(nil)
	return hash160Data

}

/*
	4.对于第二步生成的结果，使用 SHA256(SHA256(payload)) 再哈希，计算校验和。校验和是取结果哈希的前四个字节。
 */
func hashCheckSum(payloadHash []byte) []byte {
	firstHash := sha256.Sum256(payloadHash)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:CHECKSUM_BIT_LEN]
}

/**
	反向验证钱包地址合法性
	思路：将checkSum值和ripemd160_versionHash 值取出来，然后再将ripemd160_versionHash两次sha256 hash判断其结果是否等于checkSum值
 */
func VerifyWalletAddress(walletAddress []byte) bool {

	//s1: base58Decode 得到原始字节 version + publicKey + checksum
	version_pubKey_checksumBytes := crypto.Base58Decode(walletAddress)
	fmt.Println("len version_pubKey_checksumBytes = ",len(version_pubKey_checksumBytes))
	addressLen := len(version_pubKey_checksumBytes)
	//判断是否等于25byte 长度
	if addressLen != WALLET_ADDRESS_LEN {
		fmt.Println("incorrect wallet address length")
		return false
	}

	//二重验证内容Version + publicKeyHash + checkSum 合法性
	//拆分 25 = 21 + 4
	dec_ripemd160_versionHashPayloadBytes := version_pubKey_checksumBytes[:(WALLET_ADDRESS_LEN - CHECKSUM_BIT_LEN)]
	dec_checksumHashBytes := version_pubKey_checksumBytes[(WALLET_ADDRESS_LEN - CHECKSUM_BIT_LEN): ]

	//拆分后各自重新构建再判断
	rebuildPayloadChecksumBytes := hashCheckSum(dec_ripemd160_versionHashPayloadBytes)

	//判断
	if bytes.Compare(rebuildPayloadChecksumBytes, dec_checksumHashBytes) != 0 {
		fmt.Println("incorrect wallet address hash")
		return false
	}

	return true
}