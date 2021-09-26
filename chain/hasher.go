package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
)

func GetBlockHash(block Block) string {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)

	encoder.Encode(block)

	hasher := sha256.New()
	hasher.Write(buf.Bytes())
	sha := hasher.Sum(nil)

	return hex.EncodeToString(sha)
}
