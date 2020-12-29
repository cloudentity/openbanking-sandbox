package main

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"time"

	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("PORT", "8091")
	viper.SetDefault("DBFILE", "./my.db")
	viper.SetDefault("ACPURL", "")
	viper.SetDefault("ACPINTERNALURL", "")
	viper.SetDefault("APPHOST", "")
	viper.SetDefault("UIURL", "")
	viper.SetDefault("CERTFILE", "")
	viper.SetDefault("KEYFILE", "")
}

type ClientConfig struct {
	TenantID string
	ServerID string
	ClientID string
}

type LoginConfig struct {
	ClientConfig
	RootCA  string
	Timeout time.Duration
}

type HTTPClientConfig struct {
	RootCA   string
	CertFile string
	KeyFile  string
	Timeout  time.Duration
}

type AcpClient struct {
	ClientConfig
	HTTPClientConfig
}

type BankID string

type BankConfig struct {
	ID        BankID
	URL       string
	AcpClient AcpClient
}

type Config struct {
	Port           int
	DBFile         string
	ACPURL         string
	ACPInternalURL string
	AppHost        string
	UIURL          string
	CertFile       string
	KeyFile        string
	Login          LoginConfig
	Banks          []BankConfig
}

func LoadConfig() (Config, error) {
	var (
		config = Config{}
		err    error
	)

	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./data")

	if err = viper.ReadInConfig(); err != nil {
		return config, err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}

type SystemClientConfig struct {
	ClientConfig
	HTTPClientConfig
	TokenURL string
}

func (c *Config) ToSystemClientConfig(cfg BankConfig) SystemClientConfig {
	return SystemClientConfig{
		HTTPClientConfig: cfg.AcpClient.HTTPClientConfig,
		ClientConfig:     cfg.AcpClient.ClientConfig,
		TokenURL:         fmt.Sprintf("%s/%s/%s/oauth2/token", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID),
	}
}

type WebClientConfig struct {
	ClientConfig
	HTTPClientConfig
	AuthorizeURL string
	TokenURL     string
	UserinfoURL  string
	RedirectURL  string
}

func (w *WebClientConfig) GetSigningKey() (signingKey interface{}, err error) {
	var bs []byte

	if bs, err = ioutil.ReadFile(w.KeyFile); err != nil {
		return signingKey, err
	}

	block, _ := pem.Decode(bs)

	if signingKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
		return signingKey, err
	}

	return signingKey, nil
}

func (c *Config) ToWebClientConfig(cfg BankConfig) WebClientConfig {
	return WebClientConfig{
		HTTPClientConfig: cfg.AcpClient.HTTPClientConfig,
		ClientConfig:     cfg.AcpClient.ClientConfig,
		TokenURL:         fmt.Sprintf("%s/%s/%s/oauth2/token", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID),
		AuthorizeURL:     fmt.Sprintf("%s/%s/%s/oauth2/authorize", c.ACPURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID),
		UserinfoURL:      fmt.Sprintf("%s/%s/%s/userinfo", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID),
		RedirectURL:      fmt.Sprintf("%s/api/callback", c.UIURL),
	}
}

type LoginClientConfig struct {
	RootCA      string
	Timeout     time.Duration
	UserinfoURL string
}

func (c *Config) ToLoginClientConfig() LoginClientConfig {
	return LoginClientConfig{
		RootCA:      c.Login.RootCA,
		Timeout:     c.Login.Timeout,
		UserinfoURL: fmt.Sprintf("%s/%s/%s/userinfo", c.ACPInternalURL, c.Login.TenantID, c.Login.ServerID),
	}
}
