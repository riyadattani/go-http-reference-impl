package http

import (
	"github.com/quii/go-http-reference-impl/internal/ports"
	"net/http"
)

func NewWebServer(
	config ServerConfig,
	greeter ports.GreeterService,
) *http.Server {
	return &http.Server{
		Addr:         config.TCPAddress(),
		Handler:      newRouter(greeter),
		ReadTimeout:  config.HTTPReadTimeout,
		WriteTimeout: config.HTTPWriteTimeout,
	}
}
