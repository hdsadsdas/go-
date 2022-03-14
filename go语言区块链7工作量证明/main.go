package main

import (
	"fmt"
	"go语言区块链/go语言区块链7工作量证明/BLC"
)

func main(){

	// 创世区块
	 blockchain := BLC.CreateBlockchainWithGenesisBlock()

  	// 新区块
     blockchain.AddBlockToBlockchain("sand 100RMB TO haha",blockchain.Blocks[len(blockchain.Blocks)-1].Height + 1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockchain("sand 200RMB TO lala",blockchain.Blocks[len(blockchain.Blocks)-1].Height + 1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockchain("sand 300RMB TO xixi",blockchain.Blocks[len(blockchain.Blocks)-1].Height + 1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	blockchain.AddBlockToBlockchain("sand 50RMB TO hehe",blockchain.Blocks[len(blockchain.Blocks)-1].Height + 1, blockchain.Blocks[len(blockchain.Blocks)-1].Hash)


	 fmt.Println(blockchain)
	 fmt.Println(blockchain.Blocks)
	 fmt.Println(blockchain.Blocks[2])

}
