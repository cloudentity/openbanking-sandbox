package main

import (
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port        int           `env:"PORT" envDefault:"8090"`
	ClientID    string        `env:"CLIENT_ID,required"`
	TokenURL    *url.URL      `env:"TOKEN_URL,required"`
	AuthURL     *url.URL      `env:"AUTH_URL,required"`
	RedirectURL *url.URL      `env:"REDIRECT_URL,required"`
	Timeout     time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA      string        `env:"ROOT_CA,required"`
	CertFile    string        `env:"CERT_FILE,required"`
	KeyFile     string        `env:"KEY_FILE,required"`
}

func LoadConfig() (config Config, err error) {
	if err := env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}
