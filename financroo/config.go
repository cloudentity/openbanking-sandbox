package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/spf13/viper"

	acpclient "github.com/cloudentity/acp-client-go"
)

func init() {
	viper.SetDefault("PORT", "8091")
	viper.SetDefault("DB_FILE", "./data/my.db")
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

func NewAcpClient(c Config, cfg BankConfig) (acpclient.Client, error) {
	var (
		tokenURL, issuerURL, authorizeURL, userinfoURL, redirectURL *url.URL
		client                                                      acpclient.Client
		err                                                         error
	)

	if tokenURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/token", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if issuerURL, err = url.Parse(fmt.Sprintf("%s/%s/%s", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if authorizeURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/oauth2/authorize", c.ACPURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if userinfoURL, err = url.Parse(fmt.Sprintf("%s/%s/%s/userinfo", c.ACPInternalURL, cfg.AcpClient.TenantID, cfg.AcpClient.ServerID)); err != nil {
		return client, err
	}

	if redirectURL, err = url.Parse(fmt.Sprintf("%s/api/callback", c.UIURL)); err != nil {
		return client, err
	}

	config := acpclient.Config{
		ClientID:                    cfg.AcpClient.ClientID,
		IssuerURL:                   issuerURL,
		TokenURL:                    tokenURL,
		AuthorizeURL:                authorizeURL,
		UserinfoURL:                 userinfoURL,
		RedirectURL:                 redirectURL,
		RequestObjectSigningKeyFile: cfg.AcpClient.KeyFile,
		Scopes:                      []string{"accounts", "openid", "offline_access"},
		Timeout:                     cfg.AcpClient.Timeout,
		CertFile:                    cfg.AcpClient.CertFile,
		KeyFile:                     cfg.AcpClient.KeyFile,
		RootCA:                      cfg.AcpClient.RootCA,
	}

	if client, err = acpclient.New(config); err != nil {
		return client, err
	}

	return client, nil
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
