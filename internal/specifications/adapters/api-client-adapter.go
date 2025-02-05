package adapters

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIClientLogger interface {
	Log(...interface{})
	Logf(string, ...interface{})
}

type APIClient struct {
	baseURL    string
	httpClient *http.Client
	logger     APIClientLogger
}

func NewAPIClient(baseURL string, logger APIClientLogger) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
		logger: logger,
	}
}

func (a *APIClient) CheckIfHealthy() error {
	url := a.baseURL + "/internal/healthcheck"
	a.logger.Log("GET", url)

	res, err := a.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from POST %q", res.StatusCode, url)
	}

	return nil
}

func (a *APIClient) WaitForAPIToBeHealthy(retries int) error {
	var err error

	for retries > 0 {
		if err = a.CheckIfHealthy(); err != nil {
			retries -= 1
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}
	return fmt.Errorf("given up checking healthcheck, %v", err)
}

func (a *APIClient) Greet(name string) (string, error) {
	url := a.baseURL + "/greet/" + name
	a.logger.Log("GET", url)

	res, err := a.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d from GET %q", res.StatusCode, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

