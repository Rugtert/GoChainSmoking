package chain

type Block struct {
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
