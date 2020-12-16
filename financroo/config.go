package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	TenantID    string `yaml:"tenant_id"`
	ServerID    string `yaml:"server_id"`
	ClientID    string `yaml:"client_id"`
	RedirectURL string `yaml:"redirect_url"`
}

type LoginConfig struct {
	ClientConfig `yaml:",inline"`
	UIURL        string        `yaml:"ui_url"`
	RootCA       string        `yaml:"root_ca"`
	Timeout      time.Duration `yaml:"timeout"`
}

type HTTPClientConfig struct {
	RootCA   string        `yaml:"root_ca"`
	CertFile string        `yaml:"cert_file"`
	KeyFile  string        `yaml:"key_file"`
	Timeout  time.Duration `yaml:"timeout"`
}

type AcpClient struct {
	ClientConfig     `yaml:",inline"`
	HTTPClientConfig `yaml:",inline"`
}

type BankID string

type BankConfig struct {
	ID        BankID    `yaml:"id"`
	URL       string    `yaml:"url"`
	AcpClient AcpClient `yaml:"acp_client"`
}

type YAMLConfig struct {
	Login LoginConfig  `yaml:"login"`
	Banks []BankConfig `yaml:"banks"`
}

type Config struct {
	YAMLConfig
	Port           int    `env:"PORT" envDefault:"8091"`
	CertFile       string `env:"CERT_FILE,required"`
	KeyFile        string `env:"KEY_FILE,required"`
	ACPURL         string `env:"ACP_URL,required"`
	ACPInternalURL string `env:"ACP_INTERNAL_URL,required"`
	AppHost        string `env:"APP_HOST,required"`
	YAMLConfigFile string `env:"YAML_CONFIG_FILE" envDefault:"./config.yaml"`
	DBFile         string `env:"DB_FILE" envDefault:"./my.db"`
}

func LoadConfig() (Config, error) {
	var (
		config = Config{}
		bs     []byte
		err    error
	)

	if err = env.Parse(&config); err != nil {
		return config, err
	}

	if bs, err = ioutil.ReadFile(config.YAMLConfigFile); err != nil {
		return config, err
	}

	resolvedConfig := os.ExpandEnv(string(bs))

	if err = yaml.Unmarshal([]byte(resolvedConfig), &config.YAMLConfig); err != nil {
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
