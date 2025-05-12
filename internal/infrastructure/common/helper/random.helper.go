package helper

import (
	"math/rand"
	inf "restuwahyu13/shopping-cart/internal/domain/interface/helper"
	"strconv"
	"time"
)

type random struct{}

func NewRandom() inf.IRandom {
	return &random{}
}

func (h random) AlphaCharacters(length int) string {
	alphaCharacters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.NewSource(time.Now().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = alphaCharacters[rand.Intn(len(alphaCharacters))]
	}

	return string(b)
}

func (h random) Numeric(length int) string {
	rand.NewSource(time.Now().Unix())
	randOtpCode := strconv.FormatInt(int64(rand.Int()), 10)

	return randOtpCode[:length]
}
