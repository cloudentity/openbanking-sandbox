package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port        int           `env:"PORT" envDefault:"8090"`
	ClientID    string        `env:"CLIENT_ID,required"`
	IssuerURL   *url.URL      `env:"ISSUER_URL,required"`
	RedirectURL *url.URL      `env:"REDIRECT_URL,required"`
	Timeout     time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA      string        `env:"ROOT_CA,required"`
	CertFile    string        `env:"CERT_FILE,required"`
	KeyFile     string        `env:"KEY_FILE,required"`
}

func (c *Config) GetSigningKey() (signingKey interface{}, err error) {
	var bs []byte

	if bs, err = ioutil.ReadFile(c.KeyFile); err != nil {
		return signingKey, err
	}

	block, _ := pem.Decode(bs)

	if signingKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return signingKey, err
	}

	return signingKey, nil
}

func LoadConfig() (config Config, err error) {
	if err := env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}
