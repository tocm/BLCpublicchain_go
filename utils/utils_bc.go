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

/*
	比较HASH是否相等
 */
func CompareHash(hash1 []byte, hash2[] byte ) bool{

	var hashInt1 big.Int
	hashInt1.SetBytes(hash1)

	var hashInt2 big.Int
	hashInt2.SetBytes(hash2)

	// Cmp compares x and y and returns:
	//
	//   -1 if x <  y
	//    0 if x == y
	//   +1 if x >  y
	if hashInt1.Cmp(&hashInt2) == 0 {
		return true
	}
	return false
}

func ByteToString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}