package wallet

import (
	"GoChainSmoking/util"
	"bytes"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/mr-tron/base58"
)

type Wallet struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
	Address    []byte
}

var Wallets []Wallet = []Wallet{}

func FindWalletByAddress(address string) *Wallet {
	for _, w := range Wallets {
		if bytes.Equal(w.Address, []byte(address)) {
			return &w
		}
	}
	return nil
}

func FindWalletByPubkey(pubKey []byte) *Wallet {
	for _, w := range Wallets {
		if bytes.Equal(util.PublicKeyToBytes(&w.PublicKey), pubKey) {
			return &w
		}
	}
	return nil
}

func CreateWallet() *Wallet {
	fmt.Println(time.Now())

	priv, pub := util.GenerateKeyPair()
	fmt.Println(time.Now())
	base58address := base58.Encode(util.PublicKeyHash(util.PublicKeyToBytes(pub)))
	fmt.Println(time.Now())
	Wallet := Wallet{
		*priv,
		*pub,
		[]byte(base58address)}
	fmt.Println(time.Now())
	Wallets = append(Wallets, Wallet)
	fmt.Println(time.Now())
	return &Wallet
}

func (w Wallet) Base58DecodeAddress() []byte {
	res, err := base58.Decode(string(w.Address))
	util.HandleError(err)
	return res
}

func (wallet Wallet) PrintWallet() {
	fmt.Printf("\tPrivateKey: %s\n", base58.Encode(util.PrivateKeyToBytes(&wallet.PrivateKey)))
	fmt.Printf("\tPublicKey: %s\n", base58.Encode(util.PublicKeyToBytes(&wallet.PublicKey)))
	fmt.Printf("\tAddress: %s\n", wallet.Address)

}

func PrintAllWalletAddresses() {
	for _, w := range Wallets {
		fmt.Printf("\tAddress: %s\n", w.Address)
	}
}
