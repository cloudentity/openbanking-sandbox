package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("PORT", "8091")
	viper.SetDefault("DB_FILE", "./my.db")
	viper.SetDefault("ACP_URL", "")
	viper.SetDefault("ACP_INTERNAL_URL", "")
	viper.SetDefault("APP_HOST", "")
	viper.SetDefault("UI_URL", "")
	viper.SetDefault("CERT_FILE", "")
	viper.SetDefault("KEY_FILE", "")
}

type ClientConfig struct {
	TenantID string `mapstructure:"tenant_id"`
	ServerID string `mapstructure:"server_id"`
	ClientID string `mapstructure:"client_id"`
}

type LoginConfig struct {
	ClientConfig `mapstructure:",squash"`
	RootCA       string `mapstructure:"root_ca"`
	Timeout      time.Duration
}

type HTTPClientConfig struct {
	RootCA   string `mapstructure:"root_ca"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
	Timeout  time.Duration
}

type AcpClient struct {
	ClientConfig     `mapstructure:",squash"`
	HTTPClientConfig `mapstructure:",squash"`
}

type BankID string

type BankConfig struct {
	ID        BankID
	URL       string
	AcpClient AcpClient `mapstructure:"acp_client"`
}

type Config struct {
	Port           int
	DBFile         string `mapstructure:"db_file"`
	ACPURL         string `mapstructure:"acp_url" validate:"required,url"`
	ACPInternalURL string `mapstructure:"acp_internal_url" validate:"required,url"`
	AppHost        string `mapstructure:"app_host" validate:"required"`
	UIURL          string `mapstructure:"ui_url" validate:"required,url"`
	CertFile       string `mapstructure:"cert_file" validate:"required"`
	KeyFile        string `mapstructure:"key_file" validate:"required"`
	Login          LoginConfig
	Banks          []BankConfig
}

func (c *Config) Validate() (err error) {
	if err = Validator.Struct(c); err != nil {
		return err
	}

	return nil
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
