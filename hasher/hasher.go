package hasher

import (
	"GoChainSmoking/datablock"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

func GetBlockHash(block datablock.Datablock) string {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	hashableBlock := datablock.HashLessBlock{
		Data: block.Data,
		PreviousHash: block.PreviousHash,
		TimeStamp:    block.TimeStamp}
	encoder.Encode(hashableBlock)

	hasher := sha256.New()
	hasher.Write(buf.Bytes())
	sha := hasher.Sum(nil)

	return hex.EncodeToString(sha)
}
