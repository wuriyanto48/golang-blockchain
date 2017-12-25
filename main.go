package main

import (
	"fmt"
	"time"

	"github.com/wuriyanto48/golang-blockchain/core"
)

func main() {
	data1 := core.NewTransaction("Alex", "Wury", 5000000.0)
	data2 := core.NewTransaction("Wury", "Ben", 4000000.0)

	block1 := core.NewBlock(1, data1, time.Now())
	block2 := core.NewBlock(2, data2, time.Now())

	blockChain := core.NewBlockChain()

	blockChain.AddBlock(block1)
	blockChain.AddBlock(block2)
	fmt.Println(blockChain.IsChainStillValid())

	// // cheat
	// blockChain.Blocks[1].Data = core.NewTransaction("Alex", "Wury", 5500000.0)
	//
	// // will return error, because block chain is not valid anymore
	// fmt.Println(blockChain.IsChainStillValid())

	data3 := core.NewTransaction("Deny", "Alex", 999000.0)

	block3 := core.NewBlock(3, data3, time.Now())

	blockChain.AddBlock(block3)

	fmt.Println(blockChain.IsChainStillValid())

	for _, v := range blockChain.GetBlocks() {
		fmt.Println(v)
	}
}
