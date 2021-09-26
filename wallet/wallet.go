package wallet

import (
	"GoChainSmoking/util"
	"crypto/rsa"

	"github.com/mr-tron/base58"
)

type Wallet struct {
	privateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
	address    string
}

func CreateWallet() *Wallet {
	priv, pub := util.CreateKeyPair()
	base58address := base58.Encode(util.PublicKeyToBytes(&pub))

	Wallet := Wallet{
		priv,
		pub,
		base58address}
	return &Wallet
}
