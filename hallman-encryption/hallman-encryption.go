package main

import (
	"crypto/rand"
	"math/big"
)

func generatePrivateKey(primeNum *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, primeNum)
}

func generatePublicKey(privateKey, genNum, primeNum *big.Int) *big.Int {
	return new(big.Int).Exp(genNum, privateKey, primeNum)
}

func generateSharedSecret(privateKey, receivedPublicKey, primeNum *big.Int) *big.Int {
	return new(big.Int).Exp(receivedPublicKey, privateKey, primeNum)
}
