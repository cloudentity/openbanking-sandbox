package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

var Validator *validator.Validate

func main() {
	var (
		server Server
		err    error
	)

	Validator = validator.New()

	if server, err = NewServer(); err != nil {
		logrus.WithError(err).Fatalf("failed to init server")
	}

	if err = server.Start(); err != nil {
		logrus.WithError(err).Fatalf("failed to start server")
	}
}
