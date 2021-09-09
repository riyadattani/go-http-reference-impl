package http

import (
	"github.com/quii/go-http-reference-impl/internal/ports"
	"net/http"
)

func NewWebServer(
	config ServerConfig,
	greeter ports.GreeterService,
	blog ports.Blog,
) *http.Server {
	return &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      newRouter(greeter, blog),
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
