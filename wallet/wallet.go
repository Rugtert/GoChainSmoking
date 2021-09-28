package wallet

import (
	"GoChainSmoking/util"
	"crypto/rsa"

	"github.com/mr-tron/base58"
)

type Wallet struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
	Address    []byte
}

func CreateWallet() *Wallet {
	priv, pub := util.GenerateKeyPair()
	base58address := base58.Encode(util.PublicKeyToBytes(pub))

	Wallet := Wallet{
		*priv,
		*pub,
		[]byte(base58address)}
	return &Wallet
}

func (w Wallet) Base58DecodeAddress() []byte {
	res, err := base58.Decode(string(w.Address))
	util.HandleError(err)
	return res
}
