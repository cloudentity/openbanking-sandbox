package main

import (
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Client AcpClient
}

type Config struct {
	ClientID           string        `env:"CLIENT_ID"`
	ClientSecret       string        `env:"CLIENT_SECRET"`
	TokenURL           string        `env:"TOKEN_URL"`
	Timeout            time.Duration `env:"TIMEOUT"`
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

	r.GET("/", s.Get())
	r.POST("/", s.Post())

	return r.Run()
}
