package encryption

import (
	"crypto/rand"
	"crypto/rsa"
)

func generateKey() (*rsa.PrivateKey, rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// The public key is a part of the *rsa.PrivateKey struct
	publicKey := privateKey.PublicKey

	return privateKey, publicKey
}
