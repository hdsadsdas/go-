package main

import (
	"fmt"
	"go语言区块链/go语言区块链1/BLC"
)

func main(){

	block := BLC.NewBlock("haha Block",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})

	fmt.Println(block)

}
