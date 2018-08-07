package blc

type TXInput struct {
	// 1. 交易的Hash
	TxHash      []byte
	// 2. 存储TXOutput在Vout里面的索引
	VOutId      int
	// 3. 用户名
	UserKey string

}

