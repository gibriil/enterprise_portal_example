package internal

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Configuration struct {
	ServerConfig *ServerConfig
	Secrets      *Secrets
}

type ServerConfig struct {
	Address        string   `yaml:"host"`
	Port           int      `yaml:"port"`
	Protocol       string   `yaml:"protocol"`
	BaseUrl        string   `yaml:"baseUrl"`
	Environment    string   `yaml:"env"`
	TrustedOrigins []string `yaml:"trustedOrigins"`
}

type Secrets struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Application struct {
	Config            *Configuration
	Log               *slog.Logger
	Router            http.Handler
	CurrentReqContext context.Context
}

type UserContext string

func (app *Application) Location() string {
	return fmt.Sprintf("%s:%d", app.Config.ServerConfig.Address, app.Config.ServerConfig.Port)
}

func (app *Application) Serve() error {

	server := &http.Server{
		Addr:     app.Location(),
		Handler:  app.Router,
		ErrorLog: slog.NewLogLogger(app.Log.Handler(), slog.LevelError),
	}

	shutdownError := make(chan error)

	go func() {
		shutdownSignalListener := make(chan os.Signal, 1)
		signal.Notify(shutdownSignalListener, syscall.SIGINT, syscall.SIGTERM)

		signalReader := <-shutdownSignalListener

		app.Log.Info("Shutting Down Go Server", slog.String("signal", signalReader.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		shutdownError <- nil
	}()

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		app.Log.Error("Failed to Start Go Server. Could not make TCP connection", slog.String("location", server.Addr), slog.String("protocol", app.Config.ServerConfig.Protocol))
		return err
	}

	app.Log.Info(fmt.Sprintf("Starting %s Go Server", strings.ToUpper(app.Config.ServerConfig.Protocol)), slog.String("location", server.Addr+app.Config.ServerConfig.BaseUrl))
	if app.Config.ServerConfig.Protocol == "fcgi" {
		err = fcgi.Serve(listener, server.Handler)
	} else {
		err = server.Serve(listener)
	}
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.Log.Info("Go Server Stopped", slog.String("location", server.Addr))

	return nil
}
