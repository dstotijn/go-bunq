package bunq

import (
	"crypto/rsa"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

const (
	baseURL         = "https://api.bunq.com"
	apiVersion      = "v1"
	clientVersion   = "1.0.0"
	userAgent       = "go-bunq/" + clientVersion
	timestampLayout = "2006-01-02 15:04:05.000000"
)

// Client is the API client for the public bunq API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
	Token      string
	PrivateKey *rsa.PrivateKey
}

// NewClient returns a new Client.
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
	}
}

func setCommonHeaders(r *http.Request) {
	r.Header.Set("Cache-Control", "no-cache")
	r.Header.Set("User-Agent", userAgent)
	r.Header.Set("X-Bunq-Client-Request-Id", uuid.NewV4().String())
	r.Header.Set("X-Bunq-Geolocation", "0 0 0 0 NL")
	r.Header.Set("X-Bunq-Language", "en_US")
	r.Header.Set("X-Bunq-Region", "en_US")
}

func parseTimestamp(value string) (time.Time, error) {
	return time.Parse(timestampLayout, value)
}
