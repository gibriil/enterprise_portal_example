package main

import (
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		logger.Error("Failed to make tcp connection")
	}
	defer listener.Close()

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_SERVER := fcgi.ProcessEnv(r)

		data, err := json.Marshal(_SERVER)
		if err != nil {
			panic(err)
		}
		w.Write(data)
	})

	server := &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: router,
	}

	logger.Info("Starting FCGI Go Server", slog.String("location", server.Addr))

	err = fcgi.Serve(listener, server.Handler)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
