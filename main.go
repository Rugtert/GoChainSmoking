package main

import (
	"GoChainSmoking/chain"
	"GoChainSmoking/hasher"
	"fmt"
)

func main() {
	blockchain := chain.InitChain()
	fmt.Println(blockchain)

	genesis := chain.GetLastBlock(blockchain)
	fmt.Println(genesis)

	for i := 0; i < 99; i++ {
		chain.AddToChain([]string{"data1", "data2"}, blockchain)
	}
	currentBlock2 := chain.GetLastBlock(blockchain)
	fmt.Println(
		"\n ",
		"Block 2 Timestamp: ",
		currentBlock2.TimeStamp,
		"\n ",
		"Block 2 data: ",
		currentBlock2.Data,
		"\n ",
		"Block 2 Blockhash: ",
		currentBlock2.BlockHash,
		"\n ")
	fmt.Println(hasher.GetBlockHash(currentBlock2))

}
