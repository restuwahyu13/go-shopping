package pinf

import popt "restuwahyu13/shopping-cart/internal/domain/output/pkg"

type IJsonWebToken interface {
	Sign(prefix string, body any) (*popt.SignMetadata, error)
}
