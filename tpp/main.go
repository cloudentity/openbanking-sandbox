package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Config struct {
	Port         int           `env:"PORT" envDefault:"8090"`
	ClientID     string        `env:"CLIENT_ID,required"`
	AuthorizeURL *url.URL      `env:"AUTHORIZE_URL,required"`
	TokenURL     *url.URL      `env:"TOKEN_URL,required"`
	IssuerURL    *url.URL      `env:"ISSUER_URL,required"`
	UserinfoURL  *url.URL      `env:"USERINFO_URL,required"`
	RedirectURL  *url.URL      `env:"REDIRECT_URL,required"`
	Timeout      time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA       string        `env:"ROOT_CA,required"`
	CertFile     string        `env:"CERT_FILE,required"`
	KeyFile      string        `env:"KEY_FILE,required"`
	BankURL      *url.URL      `env:"BANK_URL,required"`
}

func (c *Config) ClientConfig() acpclient.Config {
	return acpclient.Config{
		ClientID:  c.ClientID,
		IssuerURL: c.IssuerURL,
		Scopes:    []string{"accounts"},
		Timeout:   c.Timeout,
		CertFile:  c.CertFile,
		KeyFile:   c.KeyFile,
		RootCA:    c.RootCA,
	}
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
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config       Config
	Client       acpclient.Client
	WebClient    AcpWebClient
	BankClient   OpenbankingClient
	SecureCookie *securecookie.SecureCookie
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.Client, err = acpclient.New(server.Config.ClientConfig()); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.WebClient, err = NewAcpWebClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp web client")
	}

	server.BankClient = NewOpenbankingClient(server.Config)

	server.SecureCookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", s.Get())
	r.POST("/login", s.Login())
	r.GET("/callback", s.Callback())

	return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
}

func main() {
	var (
		server Server
		err    error
	)

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}
