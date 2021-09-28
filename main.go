package main

import (
	"GoChainSmoking/chain"
	"GoChainSmoking/cli"
	"fmt"
	"time"

	"github.com/mr-tron/base58"
)

func main() {

	cli.Execute()
	// initwallet := chain.InitChain()

	// scndWallet := wallet.CreateWallet()

	// //genesis := chain.GetLastBlock()
	// //printBlock(*genesis, "genesis")
	// trns := []*chain.Transaction{}

	// fmt.Print("\n")
	// firstTrn := chain.CreateTransaction("wut", scndWallet.Address, initwallet)

	// trns = append(trns, &firstTrn)
	// chain.AddToChain(trns)

	// fmt.Print("\n")
	// printBlock(*chain.GetLastBlock(), "Block 1")

	// fmt.Printf("decoded firsttrn msg: %s", chain.DecodeMsg(firstTrn, *scndWallet))
}

func printBlock(block chain.Block, blockName string) {
	fmt.Printf(blockName+" hash: %x\n", block.Hash)
	fmt.Printf(blockName+" Timestamp: %s\n", time.Unix(block.TimeStamp, 0))
	fmt.Printf(blockName+" nonce: %d\n", block.Nonce)
	fmt.Printf(blockName+" Previous hash: %x\n", block.PreviousHash)
	fmt.Printf(blockName+" is valid: %t\n", block.ValidateBlock())
	fmt.Print(blockName + " Transactions: \n")
	for _, trn := range block.Transactions {
		fmt.Printf("\tID: %x\n", trn.ID)
		fmt.Printf("\tMsg: %s\n", trn.Msg)
		fmt.Printf("\tRcpt: %s\n", base58.Encode(trn.Rcpt))
		fmt.Printf("\tPubKey: %s\n", base58.Encode(trn.PubKey))
		fmt.Printf("\tSignature: %x\n", trn.Signature)
		fmt.Printf("\tVerified: %t\n", trn.Verify())
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
