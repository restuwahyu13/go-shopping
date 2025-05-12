package pinf

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"

	pdto "restuwahyu13/shopping-cart/internal/domain/dto/pkg"

	popt "restuwahyu13/shopping-cart/internal/domain/output/pkg"
)

type IJose interface {
	JweEncrypt(publicKey *rsa.PublicKey, plainText string) ([]byte, *popt.JweEncryptMetadata, error)
	JweDecrypt(privateKey *rsa.PrivateKey, cipherText []byte) (string, error)
	ImportJsonWebKey(jwkKey jwk.Key) (*popt.JwkMetadata, error)
	ExportJsonWebKey(privateKey *rsa.PrivateKey) (*popt.JwkMetadata, error)
	JwtSign(options *pdto.JwtSignOption) ([]byte, error)
	JwtVerify(prefix string, token string, redis IRedis) (*jwt.Token, error)
}
