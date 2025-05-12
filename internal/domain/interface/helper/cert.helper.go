package hinf

import (
	"crypto/rsa"
)

type ICert interface {
	GeneratePrivateKey(password []byte) (string, error)
	PrivateKeyRawToKey(privateKey []byte, password []byte) (*rsa.PrivateKey, error)
	PrivateKeyToRaw(publicKey *rsa.PrivateKey) string
	PublicKeyToRaw(publicKey *rsa.PublicKey) string
	PrivateKey(value string) error
	PublicKey(value string, raw bool) ([]byte, error)
}
