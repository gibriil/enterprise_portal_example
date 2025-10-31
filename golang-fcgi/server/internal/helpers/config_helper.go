package helpers

import (
	"log/slog"
	"os"

	"github.com/gibriil/enterprise_portal_example/internal"
	"gopkg.in/yaml.v3"
)

func OpenConfigYaml(logger *slog.Logger, path string) *internal.ServerConfig {
	file, err := os.Open(path)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var config internal.ServerConfig

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	return &config
}

func OpenSecretsYaml(logger *slog.Logger, path string) *internal.Secrets {
	file, err := os.Open(path)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	var secrets internal.Secrets

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&secrets)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	return &secrets
}
