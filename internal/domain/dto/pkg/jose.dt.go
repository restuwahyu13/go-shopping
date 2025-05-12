package pdto

import (
	"crypto/rsa"
	"time"
)

type JwtSignOption struct {
	PrivateKey *rsa.PrivateKey
	Claim      interface{}
	Kid        string
	SecretKey  string
	Iss        string
	Sub        string
	Aud        []string
	Exp        time.Time
	Nbf        float64
	Iat        time.Time
	Jti        string
}
