package utils

import (
	"bytes"
	"encoding/binary"
	"log"
	"math/big"
)

// 将int转换为字节数组
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

/**
	比较hash 是否属于创世区块
 */
func IsCompGenerateBlock(hash []byte) bool {
	var hashInt big.Int
	hashInt.SetBytes(hash)
	if big.NewInt(0).Cmp(&hashInt) == 0{
		return true
	}

	return false
}