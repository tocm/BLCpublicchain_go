package storge

import (
	"os"
	"io/ioutil"
	"log"
)

const WALLETS_ADDRESS_FILE  = "wallets.dat"

/**
	是否存在文件
 */
func IsExistFile(fileName string) bool {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return false
	}
	return true
}

/**
	读取文件
 */
func ReadFile(fileName string) []byte {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}
	return fileContent
}

/**
	写入文件
 */
func SaveFile(fileName string, content []byte) {
	// 将序列化以后的数据写入到文件，原来文件的数据会被覆盖
	err := ioutil.WriteFile(fileName, content, 0644)//0644代表文件权限
	if err != nil {
		log.Panic(err)
	}
}

