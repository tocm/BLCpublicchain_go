package blc

/**
	定义区块链结构
 */
type Blockchain struct {
	blocks [] *Block //存放所有区块数组链
}

/**
	创建创世区块链
 */
func CreateGenesisBlockChain() *Blockchain  {
	blockchain := new(Blockchain)
	//创建创世区块
	block := CreateGenesisBlock();
	blockchain.blocks = append(blockchain.blocks, block)
	return blockchain
}


/**
	新增区块到区块链上
 */
func (bchain *Blockchain) AddBlock(data []byte)  {

	//get prevHash
	chainLen := len(bchain.blocks)
	//获取链中最尾的块
	lastBlock := bchain.blocks[chainLen - 1]

	//取出上个区块的hash
	prevHash := lastBlock.Hash

	//创建新区块
	newBlock := NewBlock(data, prevHash)

	//新区块追加到链尾
	bchain.blocks = append(bchain.blocks, newBlock)

}

func (blc *Blockchain) GetBlockChain() []*Block  {
	return blc.blocks
}

