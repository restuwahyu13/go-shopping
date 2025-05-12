package cinf

type IApi interface {
	Middleware()
	Router()
	Listener()
}
