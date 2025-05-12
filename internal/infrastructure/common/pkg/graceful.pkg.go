package pkg

import (
	"crypto/rand"
	"crypto/tls"
	"net/http"
	"os"

	"github.com/ory/graceful"

	cons "restuwahyu13/shopping-cart/internal/domain/constant"
	popt "restuwahyu13/shopping-cart/internal/domain/output/pkg"
	"restuwahyu13/shopping-cart/internal/infrastructure/common/helper"
)

func Graceful(Handler func() *popt.GracefulConfig) error {
	parser := helper.NewParser()
	inboundSize, _ := parser.ToInt(os.Getenv("INBOUND_SIZE"))

	h := Handler()
	secure := true

	if _, ok := os.LookupEnv("GO_ENV"); ok && os.Getenv("GO_ENV") != "development" {
		secure = false
	}

	server := http.Server{
		Handler:        h.HANDLER,
		Addr:           ":" + h.ENV.APP.PORT,
		MaxHeaderBytes: inboundSize,
		TLSConfig: &tls.Config{
			Rand:               rand.Reader,
			InsecureSkipVerify: secure,
		},
	}

	Logrus(cons.INFO, "Server listening on port %s", h.ENV.APP.PORT)
	return graceful.Graceful(server.ListenAndServe, server.Shutdown)
}
