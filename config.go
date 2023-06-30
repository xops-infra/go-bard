package go_bard

import (
	"net/http"
)

const (
	bardAPIURLv1 = "https://api.bardapi.dev"
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL string

	HTTPClient *http.Client
}

func DefaultConfig(authToken, bardApi string) ClientConfig {
	return ClientConfig{
		authToken: authToken,
		BaseURL:   bardApi,

		HTTPClient: &http.Client{},
	}
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
