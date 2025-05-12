package pkg

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"
	pdto "restuwahyu13/shopping-cart/internal/domain/dto/pkg"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	popt "restuwahyu13/shopping-cart/internal/domain/output/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"time"
)

var (
	jso    pinf.IJose   = NewJose(context.Background())
	cipher hinf.ICrypto = helper.NewCrypto()
	cert   hinf.ICert   = helper.NewCert()
	parser hinf.IParser = helper.NewParser()
)

type jsonWebToken struct {
	env *cdto.Environtment
	rds pinf.IRedis
}

func NewJsonWebToken(env *cdto.Environtment, rds pinf.IRedis) pinf.IJsonWebToken {
	return &jsonWebToken{env: env, rds: rds}
}

func (h jsonWebToken) createSecret(prefix string, body []byte) (*popt.SecretMetadata, error) {
	secretMetadata := new(popt.SecretMetadata)
	timeNow := time.Now().Format(time.UnixDate)

	cipherTextRandom := fmt.Sprintf("%s:%s:%s:%d", prefix, string(body), timeNow, h.env.JWT.EXPIRED)
	cipherTextData := hex.EncodeToString([]byte(cipherTextRandom))

	cipherSecretKey, err := cipher.SHA512Sign(cipherTextData)
	if err != nil {
		return nil, err
	}

	cipherText, err := cipher.SHA512Sign(timeNow)
	if err != nil {
		return nil, err
	}

	cipherKey, err := cipher.AES256Encrypt(cipherSecretKey, cipherText)
	if err != nil {
		return nil, err
	}

	rsaPrivateKeyPassword := []byte(cipherKey)

	privateKey, err := cert.GeneratePrivateKey(rsaPrivateKeyPassword)
	if err != nil {
		return nil, err
	}

	secretMetadata.PrivKeyRaw = privateKey
	secretMetadata.CipherKey = cipherKey

	return secretMetadata, nil
}

func (h jsonWebToken) createSignature(prefix string, body any) (*popt.SignatureMetadata, error) {
	var (
		signatureMetadata *popt.SignatureMetadata = new(popt.SignatureMetadata)
		signatureKey      string                  = fmt.Sprintf("CREDENTIAL:%s", prefix)
		signatureField    string                  = "signature_metadata"
	)

	bodyByte, err := parser.Marshal(body)
	if err != nil {
		return nil, err
	}

	secretKey, err := h.createSecret(prefix, bodyByte)
	if err != nil {
		return nil, err
	}

	rsaPrivateKey, err := cert.PrivateKeyRawToKey([]byte(secretKey.PrivKeyRaw), []byte(secretKey.CipherKey))
	if err != nil {
		return nil, err
	}

	cipherHash512 := sha512.New()
	cipherHash512.Write(bodyByte)
	cipherHash512Body := cipherHash512.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA512, cipherHash512Body)
	if err != nil {
		return nil, err
	}

	if err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA512, cipherHash512Body, signature); err != nil {
		return nil, err
	}

	signatureOutput := hex.EncodeToString(signature)

	_, jweKey, err := jso.JweEncrypt(&rsaPrivateKey.PublicKey, signatureOutput)
	if err != nil {
		return nil, err
	}

	signatureMetadata.PrivKeyRaw = secretKey.PrivKeyRaw
	signatureMetadata.SigKey = signatureOutput
	signatureMetadata.CipherKey = secretKey.CipherKey
	signatureMetadata.JweKey = *jweKey

	signatureMetadataByte, err := parser.Marshal(signatureMetadata)
	if err != nil {
		return nil, err
	}

	jwtClaim := string(signatureMetadataByte)
	jwtExpired := time.Duration(time.Minute * time.Duration(h.env.JWT.EXPIRED))

	if err := h.rds.HSetEx(signatureKey, jwtExpired, signatureField, jwtClaim); err != nil {
		return nil, err
	}

	signatureMetadata.PrivKey = rsaPrivateKey
	return signatureMetadata, nil
}

func (h jsonWebToken) Sign(prefix string, body any) (*popt.SignMetadata, error) {
	tokenKey := fmt.Sprintf("TOKEN:%s", prefix)
	signMetadata := new(popt.SignMetadata)

	tokenExist, err := h.rds.Exists(tokenKey)
	if err != nil {
		return nil, err
	}

	if tokenExist < 1 {
		signature, err := h.createSignature(prefix, body)
		if err != nil {
			return nil, err
		}

		timestamp := time.Now().Format("2006/01/02 15:04:05")
		aud := signature.SigKey[10:20]
		iss := signature.SigKey[30:40]
		sub := signature.SigKey[50:60]
		suffix := int(math.Pow(float64(h.env.JWT.EXPIRED), float64(len(aud)+len(iss)+len(sub))))

		secretKey := fmt.Sprintf("%s:%s:%s:%s:%d", aud, iss, sub, timestamp, suffix)
		secretData := hex.EncodeToString([]byte(secretKey))

		jti, err := cipher.AES256Encrypt(secretData, prefix)
		if err != nil {
			return nil, err
		}

		duration := time.Duration(time.Minute * time.Duration(h.env.JWT.EXPIRED))
		jwtIat := time.Now().UTC().Add(-duration)
		jwtExp := time.Now().Add(duration)

		tokenPayload := new(pdto.JwtSignOption)
		tokenPayload.SecretKey = signature.CipherKey
		tokenPayload.Kid = signature.JweKey.CipherText
		tokenPayload.PrivateKey = signature.PrivKey
		tokenPayload.Aud = []string{aud}
		tokenPayload.Iss = iss
		tokenPayload.Sub = sub
		tokenPayload.Jti = jti
		tokenPayload.Iat = jwtIat
		tokenPayload.Exp = jwtExp
		tokenPayload.Claim = timestamp

		tokenData, err := jso.JwtSign(tokenPayload)
		if err != nil {
			return nil, err
		}

		if err := h.rds.SetEx(tokenKey, duration, string(tokenData)); err != nil {
			return nil, err
		}

		signMetadata.Token = string(tokenData)
		signMetadata.Expired = h.env.JWT.EXPIRED

		return signMetadata, nil
	} else {
		tokenData, err := h.rds.Get(tokenKey)
		if err != nil {
			return nil, err
		}

		signMetadata.Token = string(tokenData)
		signMetadata.Expired = h.env.JWT.EXPIRED

		return signMetadata, nil
	}
}
