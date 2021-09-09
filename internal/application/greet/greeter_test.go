package greet

import (
	"github.com/quii/go-http-reference-impl/internal/specifications"
	"testing"
)

func TestHelloGreeter(t *testing.T) {
	specifications.GreetingCriteria(t, specifications.GreetingSystemFunc(HelloGreeter))
}
