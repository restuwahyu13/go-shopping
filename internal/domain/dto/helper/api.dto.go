package hdto

type Request[T any] struct {
	Req   T
	Body  T
	Param T
	Query T
}
