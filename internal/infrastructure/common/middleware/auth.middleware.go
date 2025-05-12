package middleware

import (
	"context"
	"encoding/hex"
	"fmt"
	"math"
	"net/http"
	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	hinf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	pinf "restuwahyu13/shopping-cart/internal/domain/interface/pkg"
	hopt "restuwahyu13/shopping-cart/internal/domain/output/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/pkg"
	"strings"

	"github.com/lestrrat-go/jwx/v3/jwt"
)

func Auth(jwt_expired int, redis pinf.IRedis) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				jose   pinf.IJose    = pkg.NewJose(r.Context())
				crypto hinf.ICrypto  = helper.NewCrypto()
				res    hopt.Response = hopt.Response{}
			)

			headers := r.Header.Get("Authorization")

			if !strings.Contains(headers, "Bearer") {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Authorization is required"

				helper.Api(w, res)
				return
			}

			token := strings.Split(headers, "Bearer ")[1]

			if len(strings.Split(token, ".")) != 3 {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid token format"

				helper.Api(w, res)
				return
			}

			tokenMetadata, err := jwt.ParseRequest(r, jwt.WithHeaderKey("Authorization"), jwt.WithVerify(false))
			if err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			aud, ok := tokenMetadata.Audience()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			iss, ok := tokenMetadata.Issuer()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			sub, ok := tokenMetadata.Subject()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			jti, ok := tokenMetadata.JwtID()
			if !ok {
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			timestamp := ""
			if err := tokenMetadata.Get("timestamp", &timestamp); err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			suffix := int(math.Pow(float64(jwt_expired), float64(len(aud[0])+len(iss)+len(sub))))
			secretKey := fmt.Sprintf("%s:%s:%s:%s:%d", aud[0], iss, sub, timestamp, suffix)
			secretData := hex.EncodeToString([]byte(secretKey))

			key, err := crypto.AES256Decrypt(secretData, jti)
			if err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			if _, err = jose.JwtVerify(key, token, redis); err != nil {
				pkg.Logrus(cons.ERROR, err)
				res.StatCode = http.StatusUnauthorized
				res.ErrMsg = "Invalid access token"

				helper.Api(w, res)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", key)
			h.ServeHTTP(w, r.WithContext(ctx))

			return
		})
	}
}
