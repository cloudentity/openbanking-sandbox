package main

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	logrus "github.com/sirupsen/logrus"
)

type Config struct {
	Port     int           `env:"PORT" envDefault:"8070"`
	ClientID string        `env:"CLIENT_ID,required"`
	TokenURL *url.URL      `env:"TOKEN_URL,required"`
	Timeout  time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA   string        `env:"ROOT_CA,required"`
	CertFile string        `env:"CERT_FILE,required"`
	KeyFile  string        `env:"KEY_FILE,required"`
}

func LoadConfig() (config Config, err error) {
	if err = env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

type Server struct {
	Config          Config
	AcpClient       AcpClient
	AccountsStorage *AccountsStorage
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.AcpClient, err = NewAcpClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp client")
	}

	if server.AccountsStorage, err = InitAccountsStorage(); err != nil {
		return server, errors.Wrapf(err, "failed to init accounts storage")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()

	r.GET("/accounts", s.GetAccounts())
	r.GET("/internal/accounts/:sub", s.InternalGetAccounts())

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}

func main() {
	var (
		server Server
		err    error
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}
