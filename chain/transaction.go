package chain

import (
	"GoChainSmoking/util"
	"GoChainSmoking/wallet"
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"

	"github.com/mr-tron/base58"
)

// JUST ENCRYPT THE FUCKING MESSAGE MAN. USE THE RECIPIENTS PUBLIC KEY TO ENCRYPT AND PRIVATE KEY TO DECRYPT
// https://stackoverflow.com/questions/38612279/how-to-send-a-rsa-publickey-over-a-tcp-connection-in-go
type Transaction struct {
	ID        []byte
	Msg       []byte
	Rcpt      []byte
	Signature []byte
	PubKey    []byte
}

func (tx *Transaction) Hash() {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	tx.ID = hash[:]
}

func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	util.HandleError(err)

	return encoded.Bytes()
}

func CreateTransaction(msg string, rcpt []byte, wallet wallet.Wallet) Transaction {
	trn := Transaction{nil, []byte(msg), rcpt, nil, wallet.Base58DecodeAddress()}
	trn.Msg = trn.EncodeMsg()
	trn.Hash()
	trn.Signature = trn.Sign(wallet.PrivateKey)

	return trn
}

func (trn Transaction) Sign(priv rsa.PrivateKey) []byte {
	res, err := rsa.SignPKCS1v15(rand.Reader, &priv, crypto.SHA256, trn.ID)
	util.HandleError(err)
	return res
}

func (trn Transaction) Verify() bool {
	pub := util.BytesToPublicKey(trn.PubKey)
	err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, trn.ID, trn.Signature)
	util.HandleError(err)
	return err == nil
}

func (trn Transaction) EncodeMsg() []byte {
	pubKey, err := base58.Decode(string(trn.Rcpt))
	util.HandleError(err)
	pub := util.BytesToPublicKey(pubKey)

	res, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, []byte(trn.Msg), nil)

	util.HandleError(err)

	return res
}

func DecodeMsg(trn Transaction, wallet wallet.Wallet) []byte {
	res, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, &wallet.PrivateKey, trn.Msg, nil)
	util.HandleError(err)
	return res
}
