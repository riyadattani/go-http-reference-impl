// +build unit

package http

import (
	"github.com/quii/go-http-reference-impl/internal/application/greet"
	"github.com/quii/go-http-reference-impl/internal/ports"
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"github.com/quii/go-http-reference-impl/internal/specifications/adapters"
	"net/http/httptest"
	"testing"
)

func TestNewWebServer(t *testing.T) {

	webServer := NewWebServer(
		ServerConfig{},
		ports.GreeterServiceFunc(greet.HelloGreeter),
	)

	svr := httptest.NewServer(webServer.Handler)
	defer svr.Close()

	client := adapters.NewAPIClient(svr.URL, t)

	specifications.GreetingSpecification(t, client)
}
