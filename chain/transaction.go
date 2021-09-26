package chain

import (
	"bytes"
	"crypto/rsa"
	"encoding/gob"

	"github.com/mr-tron/base58"
)

// JUST ENCRYPT THE FUCKING MESSAGE MAN. USE THE RECIPIENTS PUBLIC KEY TO ENCRYPT AND PRIVATE KEY TO DECRYPT
// https://stackoverflow.com/questions/38612279/how-to-send-a-rsa-publickey-over-a-tcp-connection-in-go
type Transaction struct {
	ID   []byte
	msg  string
	rcpt string
}

func CreateTransaction(msg string, rcpt string) {

}

func EncodeMsg(msg string, rcpt string) {
	var pub = rsa.PublicKey
	var buf bytes.Buffer
	pubKey, err := base58.Decode(rcpt)
	dec := gob.NewDecoder(&buf)
	var pub = dec.Decode(&pubKey)
	rsa.EncryptOAEP()
	ecdsa.
}
