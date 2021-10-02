package chain

import (
	"bytes"
	"crypto/sha256"
	"math/big"
)

func (block *Block) ValidateBlock() bool {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	var intHash big.Int

	if block.PreviousHash != nil {
		if !GetBlockByHash(block.PreviousHash).ValidateBlock() {
			return false
		}
	}
	blockData := bytes.Join([][]byte{
		block.PreviousHash,
		block.HashTransactions(),
		ToHex(block.TimeStamp),
		ToHex(int64(block.Nonce)),
		ToHex(int64(Difficulty))},
		[]byte{})
	blockhash := sha256.Sum256(blockData)
	intHash.SetBytes(blockhash[:])
	return intHash.Cmp(target) == -1
}
