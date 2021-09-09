package greet

import (
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	specifications.GreetingSpecification(t, specifications.GreetingSystemFunc(HelloGreeter))
}
