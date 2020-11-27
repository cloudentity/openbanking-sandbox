package main

import (
	"net/url"
	"time"

	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Config struct {
	ClientID                                 string        `env:"CLIENT_ID,required"`
	TokenURL                                 *url.URL      `env:"TOKEN_URL,required"`
	Timeout                                  time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA                                   string        `env:"ROOT_CA"`
	CertFile                                 string        `env:"CERT_FILE,required"`
	KeyFile                                  string        `env:"KEY_FILE,required"`
	BankURL                                  *url.URL      `env:"BANK_URL,required"`
	Port                                     int           `env:"PORT" envDefault:"8085"`
	ConsentSelfServiceAuthorizationServerUrl string        `env:"CONSENT_SELF_SERVICE_AUTHORIZATION_SERVER_URL,required"`
	ConsentSelfServiceClientId               string        `env:"CONSENT_SELF_SERVICE_CLIENT_ID,required"`
	ConsentSelfServiceAuthorizationServerId  string        `env:"CONSENT_SELF_SERVICE_AUTHORIZATION_SERVER_ID,required"`
	ConsentSelfServiceTenantId               string        `env:"CONSENT_SELF_SERVICE_TENANT_ID,required"`
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Client     AcpClient
	BankClient BankClient
}

func NewServer() (Server, error) {
	var (
		config Config
		server = Server{}
		err    error
	)

	if config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.Client, err = NewAcpClient(config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	server.BankClient = NewBankClient(config)

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())

	r.GET("/consents", s.ListConsents())
	r.DELETE("/consents/*id", s.RevokeConsent())

	r.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authorizationServerURL": s.Config.ConsentSelfServiceAuthorizationServerUrl,
			"clientId":               s.Config.ConsentSelfServiceClientId,
			"authorizationServerId":  s.Config.ConsentSelfServiceAuthorizationServerId,
			"tenantId":               s.Config.ConsentSelfServiceTenantId,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("web/app/build/index.html")
	})

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
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
