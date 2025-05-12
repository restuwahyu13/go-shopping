package cdto

type Config struct {
	ENV          string `env:"GO_ENV" mapstructure:"GO_ENV"`
	PORT         string `env:"PORT" mapstructure:"PORT"`
	INBOUND_SIZE int    `env:"INBOUND_SIZE" mapstructure:"INBOUND_SIZE"`
	DSN          string `env:"PG_DSN" mapstructure:"PG_DSN"`
	CSN          string `env:"REDIS_CSN" mapstructure:"REDIS_CSN"`
	JWT_EXPIRED  int    `env:"JWT_EXPIRED" mapstructure:"JWT_EXPIRED"`
}

type (
	Application struct {
		ENV          string
		PORT         string
		INBOUND_SIZE int
	}

	Redis struct {
		URL string
	}

	Postgres struct {
		URL string
	}

	Jwt struct {
		EXPIRED int
	}
)

type Environtment struct {
	APP      *Application
	REDIS    *Redis
	POSTGRES *Postgres
	JWT      *Jwt
}
