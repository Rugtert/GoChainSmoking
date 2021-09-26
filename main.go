package main

import (
	"GoChainSmoking/chain"
	"fmt"
	"time"
)

func main() {
	blockchain := chain.InitChain()

	genesis := chain.GetLastBlock()
	printBlock(*genesis, "genesis")
	trns := []*chain.Transaction{}
	for i := 0; i < 5; i++ {
		data := chain.CreateTransaction(fmt.Sprintf("data %d", i))
		trns = append(trns, data)
	}
	chain.AddToChain(trns, blockchain)

	printBlock(*chain.GetLastBlock(blockchain), "Block 1")
	trns = []*chain.Transaction{}
	for i := 0; i < 15; i++ {
		data := chain.CreateTransaction(fmt.Sprintf("data %d", i))
		trns = append(trns, data)
	}
	chain.AddToChain(trns, blockchain)

	printBlock(*chain.GetLastBlock(blockchain), "block 2")
}

func printBlock(block chain.Block, blockName string) {
	fmt.Printf(blockName+" hash: %x\n", block.Hash)
	fmt.Printf(blockName+" Timestamp: %s\n", time.Unix(block.TimeStamp, 0))
	fmt.Printf(blockName+" nonce: %d\n", block.Nonce)
	fmt.Printf(blockName+" Previous hash: %x\n", block.PreviousHash)
	fmt.Printf(blockName+" is valid: %t\n", block.ValidateBlock())
	fmt.Print(blockName + " Transactions: \n")
	for _, trn := range block.Transactions {
		fmt.Printf("\tID: %x", trn.ID)
		fmt.Printf("\tData: %s", trn.Data)
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
