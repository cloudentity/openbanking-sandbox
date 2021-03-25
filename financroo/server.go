package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
	"gopkg.in/go-playground/validator.v9"

	acpclient "github.com/cloudentity/acp-client-go"
)

type Clients struct {
	AcpClient  acpclient.Client
	BankClient OpenbankingClient
}

func InitClients(config Config) (map[BankID]Clients, error) {
	var (
		clients      = map[BankID]Clients{}
		acpWebClient acpclient.Client
		bankClient   OpenbankingClient
		err          error
	)

	for _, bank := range config.Banks {
		if acpWebClient, err = NewAcpClient(config, bank); err != nil {
			return clients, errors.Wrapf(err, "failed to init acp web client for bank: %s", bank.ID)
		}

		if bankClient, err = NewOpenbankingClient(bank); err != nil {
			return clients, errors.Wrapf(err, "failed to init client for bank: %s", bank.ID)
		}

		clients[bank.ID] = Clients{
			AcpClient:  acpWebClient,
			BankClient: bankClient,
		}
	}

	return clients, nil
}

type Server struct {
	Config       Config
	Clients      map[BankID]Clients
	SecureCookie *securecookie.SecureCookie
	DB           *bolt.DB
	UserRepo     UserRepo
	LoginClient  acpclient.Client
	Validator    *validator.Validate
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	server.Validator = validator.New()

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if err = server.Validator.Struct(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to validate config")
	}

	if server.Clients, err = InitClients(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init clients")
	}

	if server.LoginClient, err = NewLoginClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init login client")
	}

	server.SecureCookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

	if server.DB, err = InitDB(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init db")
	}

	if server.UserRepo, err = NewUserRepo(server.DB); err != nil {
		return server, errors.Wrapf(err, "failed to init user repo")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())
	r.GET("/config.json", s.WebConfig())

	r.POST("/api/connect/:bankId", s.ConnectBank())
	r.GET("/api/callback", s.ConnectBankCallback())
	r.DELETE("/api/disconnect/:bankId", s.DisconnectBank())

	r.GET("/api/accounts", s.GetAccounts())
	r.GET("/api/transactions", s.GetTransactions())
	r.GET("/api/balances", s.GetBalances())
	r.GET("/api/banks", s.ConnectedBanks())

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"featureFlags": s.Config.FeatureFlags})
	})

	return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
}
