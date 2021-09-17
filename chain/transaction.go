package chain

import (
	"GoChainSmoking/util"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"math/rand"
)

type Transaction struct {
	ID   []byte
	Data []byte
}

func (tx Transaction) Serialize() []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	util.HandleError(err)
	return buffer.Bytes()
}

func RandomTrn() *Transaction {
	rndBytes := make([]byte, 50)
	_, err := rand.Read(rndBytes)
	util.HandleError(err)

	trn := Transaction{nil, rndBytes}
	trn.ID = trn.Hash()
	return &trn
}

func CreateTransaction(data string) *Transaction {
	trn := Transaction{nil, []byte(data)}
	trn.ID = trn.Hash()
	return &trn
}

func (trn *Transaction) Hash() []byte {
	txHash := sha256.Sum256(trn.Serialize())
	return txHash[:]
}
