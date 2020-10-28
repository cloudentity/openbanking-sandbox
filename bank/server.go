package main

import (
	"net/url"
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Client AcpClient
}

type Config struct {
	ClientID           string        `env:"CLIENT_ID,required"`
	ClientSecret       string        `env:"CLIENT_SECRET,required"`
	TokenURL           url.URL       `env:"TOKEN_URL,required"`
	Timeout            time.Duration `env:"TIMEOUT" envDefault:"5s"`
	RootCA             string        `env:"ROOT_CA"`
	InsecureSkipVerify bool          `env:"INSECURE_SKIP_VERIFY"`
}

func LoadConfig() (config Config, err error) {
	if err := env.Parse(&config); err != nil {
		return config, err
	}

	return config, err
}

func NewServer() (Server, error) {
	var (
		config Config
		server = Server{}
		err    error
	)

	if config, err = LoadConfig(); err != nil {
		return server, err
	}

	if server.Client, err = NewAcpClient(config); err != nil {
		return server, err
	}

	return server, nil
}

func (s *Server) Start() error {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", s.Get())
	r.POST("/", s.Post())
	r.GET("/test", s.Test())

	return r.Run()
}
