package blc

/**
	交易输入结构
 */
type TXInput struct {
	TxHash      []byte //对应当前交易hash
	VOutId      int //对应交易输入output中的index
	WalletAddress string //来自交易输入的钱包地址

}

