package chain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// based on https://justinmeiners.github.io/tiny-blockchain/#1:7
const Difficulty = 16

func CreateProof(b *Block) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	var blockhash [32]byte
	var inthash big.Int
	nonce := 0
	for nonce < math.MaxInt64 {
		blockData := bytes.Join(
			[][]byte{
				b.PreviousHash,
				b.HashTransactions(),
				ToHex(b.TimeStamp),
				ToHex(int64(nonce)),
				ToHex(int64(Difficulty))},
			[]byte{})
		blockhash = sha256.Sum256(blockData)

		inthash.SetBytes(blockhash[:])
		if inthash.Cmp(target) == -1 {

			b.Nonce = nonce
			b.Hash = blockhash[:]

			break
		} else {
			fmt.Printf("\r%x", blockhash)

			nonce++
		}
	}
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
