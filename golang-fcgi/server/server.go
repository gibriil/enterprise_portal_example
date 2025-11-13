package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"strconv"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		logger.Error("Failed to make tcp connection")
	}
	defer listener.Close()

	router := http.NewServeMux()

	router.HandleFunc("GET /blog-feed.go", RssFeed("blog"))

	router.HandleFunc("GET /event-feed.go", RssFeed("event"))

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

type Post struct {
	UserId float64 `json:"userId"`
	PostId float64 `json:"id"`
	Title  string  `json:"title"`
	Body   string  `json:"body"`
}

var cardTemplate = `
			<ul>
				{{range $i, $post := .}}
					{{ with $post }}
					 	<li>
							<article class="card card-%s">
								<header>{{ .Title }}</header>
								 <img src="https://picsum.photos/69{{- $i -}}/40{{- $i -}}?random=%s" />
								 <p class="description">{{ .Body }}</p>
							 </article>
						</li>
					{{ end }}
				{{end}}
			</ul>
		`

func RssFeed(feedType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		jsonData, err := io.ReadAll(response.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data []map[string]any
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 5
		}

		posts := []Post{}

		for i := range limit {
			post := Post{
				UserId: data[i]["userId"].(float64),
				PostId: data[i]["id"].(float64),
				Title:  data[i]["title"].(string),
				Body:   data[i]["body"].(string),
			}

			posts = append(posts, post)
		}

		template, err := template.New("RssFeed").Parse(fmt.Sprintf(cardTemplate, feedType, feedType))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		template.Execute(w, &posts)
	}
}
