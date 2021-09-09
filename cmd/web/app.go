package main

import (
	"github.com/quii/go-http-reference-impl/internal/adapters/http"
	"github.com/quii/go-http-reference-impl/internal/application/greet"
	"github.com/quii/go-http-reference-impl/internal/ports"
)

// App holds and creates dependencies for your application
type App struct {
	ServerConfig  http.ServerConfig
	Greeter       ports.GreeterService
}

func newApp() *App {
	config := newDefaultConfig()
	return &App{
		ServerConfig:  config,
		Greeter:       ports.GreeterServiceFunc(greet.HelloGreeter),
	}
}
