package crypto

import (
	"bytes"
	"math/big"
)

/**
	关于Base58Encode：
	1. Base58是用于Bitcoin中使用的一种独特的编码方式，主要用于产生Bitcoin的钱包地址。
	2. 相比Base64区别，Base58的种子数据去掉了6个数字："0"，字母大写"O"，字母大写"I"，和字母小写"l"，以及"+"和"/"符号。
		Base58设计目的：
	1. 避免混淆。在某些字体下，数字0和字母大写O，以及字母大写I和字母小写l会非常相似。
	2. 不使用"+"和"/"的原因是非字母或数字的字符串作为帐号较难被接受。
	3. 没有标点符号，通常不会被从中间分行。
	4. 因为58不是2的整数倍，需要不断用除法去计算，计算量比base64大。
 */



//0(零)，O(大写的 o)，I(大写的i)，l(小写的 L)，+，/
var base58EncodeKey = []byte("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func Base58Encode(data []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(data)

	base := big.NewInt(int64(len(base58EncodeKey)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58EncodeKey[mod.Int64()])
	}

	reverseBytes(result)

	for b := range data {
		if b == 0x00 {
			result = append([]byte{base58EncodeKey[0]}, result...)
		} else {
			break
		}
	}
	return result
}

// Base58转字节数组，解密
func Base58Decode(data []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for b := range data {
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := data[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(base58EncodeKey, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}


// 字节数组反转
func reverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

