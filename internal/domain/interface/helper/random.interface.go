package hinf

type IRandom interface {
	AlphaCharacters(length int) string
	Numeric(length int) string
}
