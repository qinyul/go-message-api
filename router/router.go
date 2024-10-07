package router

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qinyul/messaging-api/controller"
	"github.com/qinyul/messaging-api/middleware"
)

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Method not allowed")
	err := errors.New("method not allowed")
	http.Error(w, err.Error(), http.StatusMethodNotAllowed)
}

func NewRouter(controller controller.MessageController) http.Handler {
	router := mux.NewRouter()

	loggingMiddleware := middleware.LoggingMiddleware
	corsMiddleware := middleware.CorsMiddleware

	middlewareChain := middleware.ChainMiddlewares(
		loggingMiddleware,
		corsMiddleware,
	)

	router.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			middlewareChain(http.HandlerFunc(controller.GetMessages)).ServeHTTP(w, r)
		} else if r.Method == "POST" {
			middlewareChain(http.HandlerFunc(controller.CreateMessage)).ServeHTTP(w, r)
		} else {
			middlewareChain(http.HandlerFunc(methodNotAllowedHandler)).ServeHTTP(w, r)
		}
	})

	router.HandleFunc("/message/{id}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			middlewareChain(http.HandlerFunc(controller.GetMessageById)).ServeHTTP(w, r)
		} else if r.Method == "PATCH" {
			middlewareChain(http.HandlerFunc(controller.UpdateMessage)).ServeHTTP(w, r)
		} else if r.Method == "DELETE" {
			middlewareChain(http.HandlerFunc(controller.DeleteMessageById)).ServeHTTP(w, r)
		} else {
			middlewareChain(http.HandlerFunc(methodNotAllowedHandler)).ServeHTTP(w, r)
		}
	})

	return router
}
