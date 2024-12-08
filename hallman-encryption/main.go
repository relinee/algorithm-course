package main

import (
	"fmt"
	"math/big"
)

func main() {
	primeNum, genNum := big.NewInt(2136279841), big.NewInt(9)

	privateKeyA, err := generatePrivateKey(primeNum)
	if err != nil {
		panic(err)
	}
	privateKeyB, err := generatePrivateKey(primeNum)
	if err != nil {
		panic(err)
	}

	publicKeyA := generatePublicKey(privateKeyA, genNum, primeNum)
	publicKeyB := generatePublicKey(privateKeyB, genNum, primeNum)

	sharedSecretA := generateSharedSecret(privateKeyA, publicKeyB, primeNum)
	sharedSecretB := generateSharedSecret(privateKeyB, publicKeyA, primeNum)

	fmt.Printf("Shared Secret A: %x\n", sharedSecretA)
	fmt.Printf("Shared Secret B: %x\n", sharedSecretB)
	fmt.Printf("Is secrets equals: %v", sharedSecretA.CmpAbs(sharedSecretB) == 0)
}
