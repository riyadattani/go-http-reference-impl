package http

import (
	"github.com/gorilla/mux"
	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal"
	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal/blog_handler"
	"github.com/quii/go-http-reference-impl/internal/adapters/http/internal/greet_handler"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"net/http"
)

func newRouter(greeter ports.GreeterService, blog ports.Blog) *mux.Router {
	greetingHandler := greet_handler.NewGreetHandler(greeter)
	blogHandler := blog_handler.NewBlogHandler(blog)

	router := mux.NewRouter()
	router.HandleFunc("/internal/healthcheck", internal.HealthCheck)
	router.HandleFunc("/greet/{name}", greetingHandler.Greet)
	router.HandleFunc("/publish", blogHandler.Publish).Methods(http.MethodPost)
	router.HandleFunc("/", blogHandler.GetIndex).Methods(http.MethodGet)

	return router
}
