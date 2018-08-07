package blc

import (
	"math/big"
	"math"
	"fmt"
)

/*	工作量证明算法实现原理：
	1. 先定义一个起始值1, hash长度是64字符=32字节=256位 因此初始值为 00000000.....1
	2. 设置难度系数bit，比如：16bit 相当于 0000 0000 0000 0001 = 4个字符0001 = 2个字节
	3. 移位操作：将hash工作量值扩大： 向左移 （256-系数位）比如： 0000.....0001 ==>向左移256-16位后： 0000 0000 0000 0001 0000 0000.....
	4. 循环递增nonce拼合header信息计算hash，直到计算得到的hash<移位后的工作难度值为此

*/


//hash 256位=32字节=16进制表示64字符)
const hashBitLen  = 256
//决定工作量证明难度系数bit为单位，做移位运算系数
const targetDegreeBit  = 16 //0000 0000 0000 0001

var maxNonce int64 = math.MaxInt64 //定义nonce 最大值

type Pow struct {
	block *Block
	proof_degree *big.Int //大数据存储类型   proof_degree 用于定义区块工作量证明的难度
}


/**
	创建新的工作量证明对象，定义难度算法

	分析：
		1. big.Int对象 1
		2. 移位求难度范围
		0000 0001
		8 - 2 = 6
		0100 0000  64
		0010 0000
		0000 0000 0000 0001 0000 0000 0000 0000 0000 0000 .... 0000
 */
func CreatePow(blc *Block) *Pow  {

	//1. 创建一个初始值为1的值:256bit  000000....1
	targetDegree := big.NewInt(1)

	//2. 定义难度算法， 左移（256 - targetDegreeBit）左称后: 000100000000....
	targetDegree = targetDegree.Lsh(targetDegree, hashBitLen - targetDegreeBit)

	//创建对象
	pow := new(Pow)
	pow.block = blc
	pow.proof_degree = targetDegree;

	fmt.Printf("pow 工作量证明难度系数(bit)：%d, 难度值0x %x , int: %d\n",targetDegreeBit, pow.proof_degree, pow.proof_degree)
	fmt.Println()
	return pow;
}


/**
	pow工作量证明算法函数
 */
func (pow *Pow)RunProofOfWork() (int64, [32]byte) {

	blc := pow.block
	var nonce int64 = 0 //初始nonce 为0
	var hashBytes [32]byte //用于存放计算出的hash

	//循环递增nonce，计算hash
	for  nonce < maxNonce {
		//1. 计算hash，每次计算拼接区块头信息得到不同的字节数组，进入hash 运算
		hashBytes = blc.CalPowHash(nonce)
		fmt.Printf("\r%x",hashBytes) //注意输出 \r参数是指输出覆盖，目点是在循环中只输出一行

		//fmt.Printf("loop hashInt:%d ", hashInt)
		isValid := pow.IsValidPowHash(hashBytes)
		if isValid {
			fmt.Printf("\n nonce: %d",nonce)
			break
		}
		nonce++ //递增
	}

	fmt.Println("")
	//fmt.Println("pow hashInt: ", hashInt)
	return nonce, hashBytes
}

/**
	3. 判断pow 运算是否满足工作量证明系数值条件
	判断hashInt是否小于Block里面的 proof_degree 工作量证明难度系数的值大小

*/
func (pow *Pow)IsValidPowHash(hash [32]byte) bool  {
	var hashInt big.Int
	//[]byte 转 Int
	hashInt.SetBytes(hash[:])

	/**
	Cmp compares x and y and returns:
	-1 if x <  y
	0 if x == y
	+1 if x >  y
	 */
	return pow.proof_degree.Cmp(&hashInt) == 1
}