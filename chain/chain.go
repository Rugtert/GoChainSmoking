package chain

import (
	"container/list"
)

func InitChain() *list.List {
	var chain = list.New()
	var block = Block{}
	block.PreviousHash = []byte("")
	block.Transactions = []*Transaction{}

	chain.PushBack(block)
	return chain
}

func AddToChain(data []*Transaction, chain *list.List) *list.List {
	var block = Block{}
	block.PreviousHash = GetLastBlock(chain).PreviousHash
	block.Transactions = data

	chain.PushBack(block)
	return chain
}

func GetLastBlock(chain *list.List) Block {
	blockE := chain.Back()
	var block = blockE.Value.(Block)
	return block
}
