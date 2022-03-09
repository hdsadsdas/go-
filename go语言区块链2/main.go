package main

import (
	"fmt"
	"go语言区块链/go语言区块链2/BLC"
)

func main(){

	genesisBlock := BLC.CreatGenesisBlock("Genesis Block....")

	fmt.Println(genesisBlock)

}
