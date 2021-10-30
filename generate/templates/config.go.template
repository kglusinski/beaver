package main

import env "github.com/Netflix/go-env"

type Config struct {
	AppName string `env:"APP_NAME""`
}

func InitConfig() (Config, error) {
	cfg := Config{}

	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
