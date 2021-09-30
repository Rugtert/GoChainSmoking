package main

import "GoChainSmoking/cli"

func main() {

	cli.InteractiveCli()

	// initwallet := chain.InitChain() CHECK

	// scndWallet := wallet.CreateWallet() CHECK

	// //genesis := chain.GetLastBlock()
	// //printBlock(*genesis, "genesis")
	// trns := []*chain.Transaction{}

	// fmt.Print("\n")
	// firstTrn := chain.CreateTransaction("wut", util.PublicKeyToBytes(&scndWallet.PublicKey), initwallet)

	// trns = append(trns, &firstTrn)
	// chain.AddToChain(trns)

	// fmt.Print("\n")
	// printBlock(*chain.GetLastBlock(), "Block 1")
	// initwallet.PrintWallet()
	// fmt.Printf("decoded firsttrn msg: %s", chain.DecodeMsg(firstTrn, *scndWallet))
}
