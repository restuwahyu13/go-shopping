package hdinf

import (
	"net/http"
)

type (
	IAuthHandler interface {
		Login(rw http.ResponseWriter, r *http.Request)
	}
)
