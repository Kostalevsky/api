package main

import (
	"context"
	"project/internal/config"
	"project/internal/logger"
	"project/internal/transport"
)

func main() {
	loggerCfg := config.Logger{
		Level: "debug",
	}

	logger.InitLogger(loggerCfg)

	repoCfg := config.Repo{
		User:     "admin",
		Pass:     "admin",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgres",
		SSLMode:  "disable",
	}

	transportCfg := config.Transport{
		Host: "127.0.0.1",
		Port: "8080",
	}

	cfg := config.Config{
		Repo:      repoCfg,
		Transport: transportCfg,
	}

	tr, err := transport.NewTransport(context.TODO(), cfg)
	if err != nil {
		logger.GetLogger().Err(err).Msg("failed to create new transport")

		return
	}

	defer func() {
		if err := tr.Close(); err != nil {
			logger.GetLogger().Err(err).Msg("failed to close transport")
		}
	}()

	if err := tr.Run(); err != nil {
		logger.GetLogger().Err(err).Msg("failed to run transport")
	}

}
