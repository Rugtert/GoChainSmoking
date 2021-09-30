package chain

import (
	"GoChainSmoking/util"
	"GoChainSmoking/wallet"
	"bytes"
)

var chain []Block

func InitChain() wallet.Wallet {
	initwallet := wallet.CreateWallet()
	genesistrn := CreateTransaction("genesis", util.PublicKeyToBytes(&initwallet.PublicKey), *initwallet)

	var block = CreateBlock(append([]*Transaction{}, &genesistrn), nil)
	chain = append(chain, *block)

	return *initwallet
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

func GetBlockByHash(hash []byte) *Block {
	for _, block := range chain {
		if bytes.Equal(block.Hash, hash) {
			return &block
		}
	}
	return nil
}
