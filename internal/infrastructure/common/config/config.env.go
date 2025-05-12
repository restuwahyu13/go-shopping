package cfg

import (
	"os"
	cdto "restuwahyu13/shopping-cart/internal/domain/dto/config"

	"github.com/caarlos0/env"

	"github.com/spf13/viper"
)

func NewEnvirontment(name, path, ext string) (*cdto.Environtment, error) {
	cfg := cdto.Config{}

	if _, ok := os.LookupEnv("GO_ENV"); !ok {
		viper.SetConfigName(name)
		viper.SetConfigType(ext)
		viper.AddConfigPath(path)
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
	} else {
		if err := env.Parse(&cfg); err != nil {
			return nil, err
		}
	}

	return &cdto.Environtment{
		APP: &cdto.Application{
			ENV:          cfg.ENV,
			PORT:         cfg.PORT,
			INBOUND_SIZE: cfg.INBOUND_SIZE,
		},
		REDIS: &cdto.Redis{
			URL: cfg.CSN,
		},
		POSTGRES: &cdto.Postgres{
			URL: cfg.DSN,
		},
		JWT: &cdto.Jwt{
			EXPIRED: cfg.JWT_EXPIRED,
		},
	}, nil
}
