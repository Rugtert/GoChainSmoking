package chain

import (
	"GoChainSmoking/datablock"
	"GoChainSmoking/hasher"
	"container/list"
	"time"
)



func InitChain() *list.List {
	var chain = list.New()
	var block = datablock.Datablock{}
	block.PreviousHash = ""
	block.Data = []string{}
	block.TimeStamp = time.Now().Unix()
	block.BlockHash = hasher.GetBlockHash(block)
	chain.PushBack(block)
	return chain
}


func AddToChain(data []string, chain *list.List) *list.List {
	var block = datablock.Datablock{}
	block.PreviousHash = GetLastBlock(chain).BlockHash
	block.Data = data
	block.TimeStamp = time.Now().Unix()
	block.BlockHash = hasher.GetBlockHash(block)
	chain.PushBack(block)
	return chain
}

func GetLastBlock(chain *list.List) datablock.Datablock {
	blockE := chain.Back()
	var block = blockE.Value.(datablock.Datablock)
	return block
}