package pkg

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"reflect"
	pdto "restuwahyu13/shopping-cart/internal/domain/dto/pkg"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	popt "restuwahyu13/shopping-cart/internal/domain/output/pkg"

	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwe"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jws"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

type jose struct {
	ctx context.Context
}

func NewJose(ctx context.Context) pinf.IJose {
	jwk.Configure(jwk.WithStrictKeyUsage(true))
	return &jose{ctx: ctx}
}

func (p jose) JweEncrypt(publicKey *rsa.PublicKey, plainText string) ([]byte, *popt.JweEncryptMetadata, error) {
	jweEncryptMetadata := new(popt.JweEncryptMetadata)

	headers := jwe.NewHeaders()
	headers.Set("sig", plainText)
	headers.Set("alg", jwa.RSA_OAEP_512().String())
	headers.Set("enc", jwa.A256GCM().String())

	cipherText, err := jwe.Encrypt([]byte(plainText), jwe.WithKey(jwa.RSA_OAEP_512(), publicKey), jwe.WithContentEncryption(jwa.A256GCM()), jwe.WithCompact(), jwe.WithJSON(), jwe.WithProtectedHeaders(headers))
	if err != nil {
		return nil, nil, err
	}

	if err := parser.Unmarshal(cipherText, jweEncryptMetadata); err != nil {
		return nil, nil, err
	}

	return cipherText, jweEncryptMetadata, nil
}

func (p jose) JweDecrypt(privateKey *rsa.PrivateKey, cipherText []byte) (string, error) {
	jwtKey, err := jwk.Import(privateKey)
	if err != nil {
		return "", err
	}

	jwkSet := jwk.NewSet()
	if err := jwkSet.AddKey(jwtKey); err != nil {
		return "", err
	}

	plainText, err := jwe.Decrypt(cipherText, jwe.WithKey(jwa.RSA_OAEP_512(), jwtKey), jwe.WithKeySet(jwkSet, jwe.WithRequireKid(false)))
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}

func (p jose) ImportJsonWebKey(jwkKey jwk.Key) (*popt.JwkMetadata, error) {
	jwkRawMetadata := popt.JwkMetadata{}

	if _, err := jwk.IsPrivateKey(jwkKey); err != nil {
		return nil, err
	}

	if err := jwk.AssignKeyID(jwkKey); err != nil {
		return nil, err
	}

	jwkKeyByte, err := parser.Marshal(&jwkKey)
	if err != nil {
		return nil, err
	}

	jwkRaw, err := jwk.ParseKey(jwkKeyByte)
	if err != nil {
		return nil, err
	}

	if err := parser.Unmarshal(jwkKeyByte, &jwkRawMetadata.KeyRaw); err != nil {
		return nil, err
	}

	jwkRawMetadata.Key = jwkRaw

	return &jwkRawMetadata, nil
}

func (p jose) ExportJsonWebKey(privateKey *rsa.PrivateKey) (*popt.JwkMetadata, error) {
	jwkRawMetadata := popt.JwkMetadata{}

	jwkRaw, err := jwk.ParseKey([]byte(cert.PrivateKeyToRaw(privateKey)), jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}

	jwkRawByte, err := parser.Marshal(&jwkRaw)
	if err != nil {
		return nil, err
	}

	if err := parser.Unmarshal(jwkRawByte, &jwkRawMetadata.KeyRaw); err != nil {
		return nil, err
	}

	jwkRawMetadata.Key = jwkRaw.(jwk.Key)

	return &jwkRawMetadata, nil
}

func (p jose) JwtSign(options *pdto.JwtSignOption) ([]byte, error) {
	jwsHeader := jws.NewHeaders()
	jwsHeader.Set("alg", jwa.RS512)
	jwsHeader.Set("typ", "JWT")
	jwsHeader.Set("cty", "JWT")
	jwsHeader.Set("kid", options.Kid)
	jwsHeader.Set("b64", true)

	jwtBuilder := jwt.NewBuilder()
	jwtBuilder.Audience(options.Aud)
	jwtBuilder.Issuer(options.Iss)
	jwtBuilder.Subject(options.Sub)
	jwtBuilder.IssuedAt(options.Iat)
	jwtBuilder.Expiration(options.Exp)
	jwtBuilder.JwtID(options.Jti)
	jwtBuilder.Claim("timestamp", options.Claim)

	jwtToken, err := jwtBuilder.Build()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Sign(jwtToken, jwt.WithKey(jwa.RS512(), options.PrivateKey, jws.WithProtectedHeaders(jwsHeader)))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (p jose) JwtVerify(prefix string, token string, redis pinf.IRedis) (*jwt.Token, error) {
	signatureKey := fmt.Sprintf("CREDENTIAL:%s", prefix)
	signatureMetadataField := "signature_metadata"

	signatureMetadata := new(popt.SignatureMetadata)
	signatureMetadataBytes, err := redis.HGet(signatureKey, signatureMetadataField)
	if err != nil {
		return nil, err
	}

	if err := parser.Unmarshal(signatureMetadataBytes, signatureMetadata); err != nil {
		return nil, err
	}

	if reflect.DeepEqual(signatureMetadata, popt.SignatureMetadata{}) {
		return nil, errors.New("Invalid secretkey or signature")
	}

	privateKey, err := cert.PrivateKeyRawToKey([]byte(signatureMetadata.PrivKeyRaw), []byte(signatureMetadata.CipherKey))
	if err != nil {
		return nil, err
	}

	exportJws, err := jws.ParseString(token)
	if err != nil {
		return nil, err
	}

	signatures := exportJws.Signatures()
	if len(signatures) < 1 {
		return nil, errors.New("Invalid signature")
	}

	jwsSignature := new(jws.Signature)
	for _, signature := range signatures {
		jwsSignature = signature
		break
	}

	jwsHeaders := jwsSignature.ProtectedHeaders()

	algorithm, ok := jwsHeaders.Algorithm()
	if !ok {
		return nil, errors.New("Invalid algorithm")
	} else if algorithm != jwa.RS512() {
		return nil, errors.New("Invalid algorithm")
	}

	kid, ok := jwsHeaders.KeyID()
	if !ok {
		return nil, errors.New("Invalid keyid")
	} else if kid != signatureMetadata.JweKey.CipherText {
		return nil, errors.New("Invalid keyid")
	}

	aud := signatureMetadata.SigKey[10:20]
	iss := signatureMetadata.SigKey[30:40]
	sub := signatureMetadata.SigKey[50:60]
	claim := "timestamp"

	jwkKey, err := jwk.Import(privateKey)
	if err != nil {
		return nil, err
	}

	_, err = jws.Verify([]byte(token), jws.WithValidateKey(true), jws.WithKey(algorithm, jwkKey), jws.WithMessage(exportJws))
	if err != nil {
		return nil, err
	}

	jwtParse, err := jwt.Parse([]byte(token),
		jwt.WithKey(algorithm, privateKey),
		jwt.WithAudience(aud),
		jwt.WithIssuer(iss),
		jwt.WithSubject(sub),
		jwt.WithRequiredClaim(claim),
	)

	if err != nil {
		return nil, err
	}

	return &jwtParse, nil
}
