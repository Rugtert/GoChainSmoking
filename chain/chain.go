package chain

import "GoChainSmoking/wallet"

var chain []Block

func InitChain() ([]Block, wallet.Wallet) {
	initwallet := wallet.CreateWallet()
	genesistrn := CreateTransaction("genesis", initwallet, initwallet.PublicKey)

	var block = CreateBlock(append([]*Transaction{}, genesistrn), nil)
	chain = append(chain, *block)
}

func AddToChain(data []*Transaction) {
	block := CreateBlock(data, GetLastBlock().Hash)
	chain = append(chain, *block)
}

func GetLastBlock() *Block {
	return &chain[len(chain)-1]
}

func GetBlockAt(index int) *Block {
	return &chain[index]
}

func GetChain() []Block {
	return chain
}
