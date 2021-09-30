package cli

import (
	"GoChainSmoking/chain"
	"GoChainSmoking/util"
	"GoChainSmoking/wallet"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/goombaio/namegenerator"
	"github.com/manifoldco/promptui"
)

func InteractiveCli() {
	for {
		prompt := promptui.Select{
			Label: "Choose command",
			Items: []string{"Initialize chain", "Create wallet", "Create wallets", "Print all wallets", "Print last block", "get my trns", "transfer", "print chain", "generate random transactions"},
		}

		_, result, err := prompt.Run()
		util.HandleError(err)

		switch result {
		case "Initialize chain":
			initChain()
		case "Create wallet":
			CreateWallet()
		case "Create wallets":
			createWallets()
		case "Print all wallets":
			printAllWallets()
		case "print chain":
			printChain()
		case "Print last block":
			printLastBlock()
		case "transfer":
			rcpt := promptui.Prompt{Label: "To what address?"}
			src := promptui.Prompt{Label: "What is your address?"}
			msg := promptui.Prompt{Label: "What message?"}
			rcptAddress, err := rcpt.Run()
			util.HandleError(err)
			srcAddress, err := src.Run()
			util.HandleError(err)
			message, err := msg.Run()
			util.HandleError(err)
			transfer(message, rcptAddress, srcAddress)
		case "get my trns":
			address := promptui.Prompt{Label: "What is your address?"}
			res, err := address.Run()
			util.HandleError(err)
			findMyTrns(res)
		case "generate random transactions":
			generateRandomTransactions(rand.Intn(10))
		default:
			fmt.Println("You choose poorly")
		}
	}
}

func generateRandomTransactions(amount int) {
	walletCount := len(wallet.Wallets)
	if walletCount < 2 {
		fmt.Println("Create more wallets first.")
		return
	}
	trns := []*chain.Transaction{}
	for i := 0; i < amount; i++ {
		seed := time.Now().UTC().UnixNano()
		rand.Seed(seed)
		firstRndWalNum := rand.Intn((walletCount - 1) + 1)
		var scnRndWalNum int

	secondgenerator:
		for {
			seed := time.Now().UTC().UnixNano()
			rand.Seed(seed)
			scnRndWalNum = rand.Intn((walletCount - 1) + 1)
			if scnRndWalNum != firstRndWalNum {
				break secondgenerator
			}
		}
		trn := chain.CreateTransaction(namegenerator.NewNameGenerator(seed).Generate(), util.PublicKeyToBytes(&wallet.Wallets[firstRndWalNum].PublicKey), wallet.Wallets[scnRndWalNum])
		trns = append(trns, &trn)
	}
	chain.AddToChain(trns)
}

func printChain() {
	chain := chain.GetChain()
	for i, block := range chain {
		blocknum := i
		block.PrintBlock(fmt.Sprintf("Block %d", blocknum))
	}
}

func initChain() wallet.Wallet {
	var initwallet = chain.InitChain()
	fmt.Printf("\tAddress: %s\n", initwallet.Address)
	return *&initwallet
}

func CreateWallet() wallet.Wallet {
	w := wallet.CreateWallet()
	fmt.Printf("\tAddress: %s\n", w.Address)
	return *w
}
func createWallets() {
	for i := 0; i < 10; i++ {
		CreateWallet()
	}
}
func printAllWallets() {
	fmt.Println("")
	wallet.PrintAllWalletAddresses()
}

func printLastBlock() {
	fmt.Print("-----------------------\n")
	chain.GetLastBlock().PrintBlock("Last block")
	fmt.Print("-----------------------\n")
}

func findMyTrns(address string) {
	bchain := chain.GetChain()
	wallet := wallet.FindWalletByAddress(address)
	if wallet == nil {
		fmt.Print("-----------------------\n")
		log.Print("wallet not found")
		fmt.Print("-----------------------\n")
		return
	}
	pubkey := util.PublicKeyToBytes(&wallet.PublicKey)
	for _, block := range bchain {
		for _, trn := range block.Transactions {
			if bytes.Equal(trn.Rcpt, pubkey) {
				fmt.Print("-----------------------\n")
				fmt.Printf("trn message: %s\n", chain.DecodeMsg(*trn, *wallet))
				fmt.Printf("trn signature: %x\n", trn.Signature)
				fmt.Printf("trn enc msg: %x\n", trn.Msg)
				fmt.Printf("trn valid: %t\n", trn.Verify())
				fmt.Print("-----------------------\n")
			}
		}
	}
}

func transfer(msg string, rcptAddress string, yourAddress string) {
	rcptwallet := wallet.FindWalletByAddress(rcptAddress)
	srcwallet := wallet.FindWalletByAddress(yourAddress)
	trn := chain.CreateTransaction(msg, util.PublicKeyToBytes(&rcptwallet.PublicKey), *srcwallet)
	chain.AddToChain(append([]*chain.Transaction{}, &trn))
}
