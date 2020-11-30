package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	SystemClientID              string        `env:"SYSTEM_CLIENT_ID,required"`
	SystemClientSecret          string        `env:"SYSTEM_CLIENT_SECRET,required"`
	SystemTokenURL              *url.URL      `env:"SYSTEM_TOKEN_URL,required"`
	SystemClientsServerID       string        `env:"SYSTEM_CLIENTS_SERVER_ID,required"`
	Timeout                     time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA                      string        `env:"ROOT_CA"`
	CertFile                    string        `env:"CERT_FILE,required"`
	KeyFile                     string        `env:"KEY_FILE,required"`
	BankURL                     *url.URL      `env:"BANK_URL,required"`
	Port                        int           `env:"PORT" envDefault:"8086"`
	LoginAuthorizationServerURL string        `env:"LOGIN_AUTHORIZATION_SERVER_URL,required"`
	LoginClientID               string        `env:"LOGIN_CLIENT_ID,required"`
	LoginAuthorizationServerID  string        `env:"LOGIN_AUTHORIZATION_SERVER_ID,required"`
	LoginTenantID               string        `env:"LOGIN_TENANT_ID,required"`
	IntrospectClientID          string        `env:"INTROSPECT_CLIENT_ID,required"`
	IntrospectClientSecret      string        `env:"INTROSPECT_CLIENT_SECRET,required"`
	IntrospectTokenURL          *url.URL      `env:"INTROSPECT_TOKEN_URL,required"`
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Client           AcpClient
	IntrospectClient AcpIntrospectClient
	BankClient       BankClient
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

	if server.IntrospectClient, err = NewAcpIntrospectClient(config); err != nil {
		return server, errors.Wrapf(err, "failed to init introspect acp client")
	}

	server.BankClient = NewBankClient(config)

	return server, nil
}

func (s *Server) Start() error {
	var (
		config Config
		err    error
	)

	r := gin.Default()

	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())

	r.GET("/clients", s.ListClients())
	r.DELETE("/consents/:id", s.RevokeConsent())

	if config, err = LoadConfig(); err != nil {
		return errors.Wrapf(err, "failed to load config")
	}

	r.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authorizationServerURL": config.LoginAuthorizationServerURL,
			"clientId":               config.LoginClientID,
			"authorizationServerId":  config.LoginAuthorizationServerID,
			"tenantId":               config.LoginTenantID,
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.File("web/app/build/index.html")
	})

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(config.Port)))
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