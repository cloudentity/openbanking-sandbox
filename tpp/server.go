package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/pkg/errors"
)

type Server struct {
	Config       Config
	Client       AcpClient
	WebClient    AcpWebClient
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

	if server.Client, err = NewAcpMTLSClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp mtls client")
	}

	if server.WebClient, err = NewAcpWebClient(server.Config); err != nil {
		return server, errors.Wrapf(err, "failed to init acp web client")
	}

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

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}
