package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Server struct {
	Config Config
}

func NewServer() (Server, error) {
	var (
		server = Server{}
		err    error
	)

	if server.Config, err = LoadConfig(); err != nil {
		return server, errors.Wrapf(err, "failed to load config")
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("web/app/build/index.html")
	r.Static("/static", "./web/app/build/static")

	r.GET("/", s.Index())
	r.POST("/login", s.Login())
	r.GET("/callback", s.Callback())

	r.GET("/config.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/accounts", s.Accounts())

	return r.Run(fmt.Sprintf(":%s", strconv.Itoa(s.Config.Port)))
}
