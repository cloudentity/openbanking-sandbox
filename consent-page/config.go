package main

import (
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ClientID           string        `env:"CLIENT_ID,required"`
	ClientSecret       string        `env:"CLIENT_SECRET,required"`
	IssuerURL          *url.URL      `env:"ISSUER_URL,required"`
	Timeout            time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA             string        `env:"ROOT_CA"`
	InsecureSkipVerify bool          `env:"INSECURE_SKIP_VERIFY"`
}

func LoadConfig() (config Config, err error) {
	if err := env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}
