package blc

import (
	"time"
	"bytes"
	"crypto/sha256"
	"BLCpublicchain_go/utils"
	"fmt"
	"encoding/gob"
	"log"
)

/**
	定义区块结构
 */
type Block struct {
	Timestampe int64
	PrevHash [32]byte //上一个block Hash 固定为32字节长度
	Hash [32]byte//当前block hash
	Data []byte
	Nonce int64

}

/**
	Create new block
 */
func NewBlock(data []byte, preHash [32]byte)  *Block {
	new_block := new(Block)
	new_block.Data = data
	new_block.Timestampe = time.Now().Unix()
	new_block.PrevHash = preHash

	//执行工作量证明 pow算法 挖矿
	pow := CreatePow(new_block)
	nonce, hash := pow.RunProofOfWork()

	new_block.Nonce = nonce
	new_block.Hash = hash

	fmt.Println("")

	return new_block
}

/**
	创建创世区块
 */
func CreateGenesisBlock() *Block {
	genesisBlock := NewBlock([]byte("This is a genesis block"), [32]byte{0} )
	return genesisBlock;
}

/**
	计算 Hash值
	原理是：通过把区块链的头信息的所有数据拼接起来，然后通过sha256算法计算出一串32个字节长度的hash值（输出显示以16进制为单位）
 */
func (blc *Block) CalPowHash(nonce int64) [32]byte {

	//取出时间需要转成byte[] 数组拼接
	timestampe := utils.IntToHex(blc.Timestampe)
	//取出preHash，由于原来定义是32byte固定数组，所以不能直接从固定数组赋值非固定长度数据，只能先执行拷贝到一个新数组
	preHash := blc.PrevHash[:]

	//拼接所有数据到一个二维数组 组成一个头信息
	header := bytes.Join([][]byte{timestampe,preHash,blc.Data,utils.IntToHex(nonce)}, []byte{})

	//生成32byte hash
	hash32 := sha256.Sum256(header)

	return hash32;

}


// 将结构序列化成字节数组
func (block *Block) EnSerialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化
func DeSerialize(blockBytes []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

func (block *Block) ShowBlockInfo()  {
	fmt.Println("prev hash", block.PrevHash)
	fmt.Println("hash: ",block.Hash)
	fmt.Println("data: ", string(block.Data))
	fmt.Println("timestamp", block.Timestampe)
	fmt.Println("nonce", block.Nonce)
	fmt.Println()
}
 
 