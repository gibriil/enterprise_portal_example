package main

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gibriil/enterprise_portal_example/internal"
	"github.com/gibriil/enterprise_portal_example/internal/helpers"
	"github.com/gibriil/enterprise_portal_example/internal/router"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	conf, err := filepath.Abs("etc/conf.d/config.yaml")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	// secrets, err := filepath.Abs("etc/conf.d/secrets.yaml")
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	os.Exit(1)
	// }

	config := &internal.Configuration{
		ServerConfig: helpers.OpenConfigYaml(logger, conf),
		// Secrets:      helpers.OpenSecretsYaml(logger, secrets),
	}

	app := &internal.Application{
		Config: config,
		Log:    logger,
	}

	app.Router = router.CreateRouter(app)

	err = app.Serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
