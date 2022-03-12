package BLC

import (
	"time"
)

//区块结构
type Block struct {
	//1，区块高度
	Height int64
	//2，上个区块的HASH
	PrevBlockHash []byte
	//3，交易数据
	Data []byte
	//4，时间戳
	Timestamp int64
	//5，Hash
	Hash []byte
	// 6 . Nonce
	Nonce int64
}



//1，创建新的区块
func NewBlock(data string,height int64,ptrvBlockHash []byte) *Block {

	//创建区块
	block := &Block{height,ptrvBlockHash,[]byte(data),time.Now().Unix(),nil,0}

	//调用工作量证明方法并且返回有效的Hash和Nonce

	pow := NewProofOfWork(block)


	hash,nonce := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block

}

//2 .单独写个方法，生成创世区块

func CreatGenesisBlock(data string)*Block{

	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

}




























