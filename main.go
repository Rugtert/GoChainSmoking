package main

import (
	"GoChainSmoking/chain"
	"GoChainSmoking/hasher"
	"fmt"
)

func main() {
	blockchain := chain.InitChain();
	fmt.Println(blockchain)

	genesis := chain.GetLastBlock(blockchain)
	fmt.Println(genesis);

	chain.AddToChain([]string{"data1","data2"},blockchain)
	currentBlock := chain.GetLastBlock(blockchain)
	fmt.Println(currentBlock);
	fmt.Println(hasher.GetBlockHash(currentBlock))

	chain.AddToChain([]string{"data3","data4"},blockchain)
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
		"\n ");
	fmt.Println(hasher.GetBlockHash(currentBlock2))

}
