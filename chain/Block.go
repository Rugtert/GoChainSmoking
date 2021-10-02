package chain

import (
	"GoChainSmoking/wallet"
	"fmt"
	"time"
)

type Block struct {
	TimeStamp    int64
	Hash         []byte
	Transactions []*Transaction
	PreviousHash []byte
	Nonce        int
}

func (block *Block) HashTransactions() []byte {
	var Hashes [][]byte
	for _, trn := range block.Transactions {
		Hashes = append(Hashes, trn.Serialize())
	}
	tree := NewMerkleTree(Hashes)
	return tree.RootNode.Data
}

func CreateBlock(trns []*Transaction, previousHash []byte) *Block {
	block := &Block{
		TimeStamp:    time.Now().Unix(),
		PreviousHash: previousHash,
		Transactions: trns}
	CreateProof(block)
	return block
}

func (block Block) PrintBlock(blockName string) {
	fmt.Print("----------------------------------------------------------")
	fmt.Printf(blockName+" hash: %x\n", block.Hash)
	fmt.Printf(blockName+" Timestamp: %s\n", time.Unix(block.TimeStamp, 0))
	fmt.Printf(blockName+" nonce: %d\n", block.Nonce)
	fmt.Printf(blockName+" Previous hash: %x\n", block.PreviousHash)
	fmt.Printf(blockName+" is valid: %t\n", block.ValidateBlock())
	fmt.Printf(blockName+" trns : %d\n", len(block.Transactions))
	fmt.Print(blockName + " Transactions: \n")
	for _, trn := range block.Transactions {
		fmt.Printf("\tID: %x\n", trn.ID)
		// fmt.Printf("\tEncrypted Msg: %s\n", trn.Msg)
		fmt.Printf("\tRcpt: %s\n", wallet.FindWalletByPubkey(trn.Rcpt).Address)
		// fmt.Printf("\tPubKey: %s\n", trn.PubKey)
		// fmt.Printf("\tSignature: %x\n", trn.Signature)
		fmt.Printf("\tVerified: %t\n", trn.Verify())
		fmt.Print("----------------------------------------------------------")
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
