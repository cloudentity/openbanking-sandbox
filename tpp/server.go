package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Config Config
	// Client AcpClient
}

func NewServer() (Server, error) {
	var (
		config Config
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	logrus.Infof("XXX config: %+v", config)

	// if server.Client, err = NewAcpClient(config); err != nil {
	// 	return server, errors.Wrapf(err, "failed to init acp client")
	// }

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
