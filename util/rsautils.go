package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

const bitSize = 4096

func PublicKeyToBytes(pub *rsa.PublicKey) []byte {
	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		HandleError(err)
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return pubBytes
}

func CreateKeyPair() (rsa.PrivateKey, rsa.PublicKey) {
	priv, err := rsa.GenerateKey(rand.Reader, bitSize)
	HandleError(err)

	return *priv, priv.PublicKey
}
