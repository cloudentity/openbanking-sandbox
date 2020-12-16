package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

type Server struct {
	Config              Config
	AccountAccessClient AcpAccountAccessClient
	WebClient           AcpWebClient
	BankClient          OpenbankingClient
	SecureCookie        *securecookie.SecureCookie
	DB                  *bolt.DB
	UserRepo            UserRepo
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	if server.AccountAccessClient, err = NewAcpAccountAccessClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp mtls client")
	}

	if server.WebClient, err = NewAcpWebClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp web client")
	}

	server.BankClient = NewOpenbankingClient(server.Config)

	server.SecureCookie = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

	if server.DB, err = InitDB(); err != nil {
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
	r.POST("/api/connect", s.Connect())
	r.GET("/api/callback", s.Callback())

	r.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"authorizationServerURL": s.Config.FinancrooAuthorizationServerURL,
			"clientId":               s.Config.FinancrooClientID,
			"authorizationServerId":  s.Config.FinancrooAuthorizationServerID,
			"tenantId":               s.Config.FinancrooTenantID,
		})
	})

	r.GET("/api/accounts", s.GetAccounts())
	r.GET("/api/transactions", s.GetTransactions())
	r.GET("/api/balances", s.GetBalances())

	r.NoRoute(func(c *gin.Context) {
		c.File("web/app/build/index.html")
	})

	return r.RunTLS(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)), s.Config.CertFile, s.Config.KeyFile)
}
