package chain

import (
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
