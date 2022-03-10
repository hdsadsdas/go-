package main

import (
	"fmt"
	"go语言区块链/go语言区块链3/BLC"
)

func main(){

	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockchain)
	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])

}
