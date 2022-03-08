package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

//获取Hash值
func (block *Block) setHash(){

	//1. 将 Height  转化为 []byte

	heightBytes := IntToHex(block.Height)

	//2. Timestamp  转化为 []byte

	timeStamp := strconv.FormatInt(block.Timestamp,2)
	timeBytes := []byte(timeStamp)

	//3 .将所有的属性拼接起来

	blockBytes := bytes.Join([][]byte{heightBytes,block.PrevBlockHash,block.Data,timeBytes,block.Hash},[]byte{})

	//4 . 生成Hash

	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]

}


//1，创建新的区块
func NewBlock(data string,height int64,ptrvBlockHash []byte) *Block {

	//创建区块
	block := &Block{height,ptrvBlockHash,[]byte(data),time.Now().Unix(),nil}

	//生成本身Hash值b
    block.setHash()

	return block

}