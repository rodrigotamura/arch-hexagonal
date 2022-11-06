package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni" // middleware
	"github.com/gorilla/mux"         // works with routes
	"github.com/rodrigotamura/arch-hexagonal/adapters/web/handler"
	"github.com/rodrigotamura/arch-hexagonal/application"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
