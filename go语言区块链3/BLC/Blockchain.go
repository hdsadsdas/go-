package BLC

//Blockchain 区块链
type Blockchain struct {

	Blocks []*Block    //存储有序的区块

}


//1 . 创建带有创世区块的区块链

func CreateBlockchainWithGenesisBlock() *Blockchain {

	//创建创世区块
  genesisBlock := CreatGenesisBlock("Genesis Data....")
   // 返回区块链对象
  return &Blockchain{[]*Block{genesisBlock}}

}