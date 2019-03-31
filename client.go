package godaddy

import (
	"net/http"

	"github.com/adrianosela/godaddy/v1/domains/records"
)

// Client holds client config
type Client struct {
	apiKey     string
	apiSecret  string
	Host       string
	HTTPClient *http.Client
}

const (
	// HostOTE is the hostname for the OTE GoDaddy API environment
	HostOTE = "https://api.ote-godaddy.com"
	// HostProd is the hostname for the Production GoDaddy API environment
	HostProd = "https://api.godaddy.com"
)

// NewClient is the Client constructor
func NewClient(key, secret, host string) *Client {
	return &Client{
		apiKey:     key,
		apiSecret:  secret,
		Host:       host,
		HTTPClient: http.DefaultClient, // fixme
	}
}

// GetRecords returns records for a given domain
func (c *Client) GetRecords(domain, recordType, name string) error {
	conf := &records.Config{
		APIKey:     c.apiKey,
		APISecret:  c.apiSecret,
		APIHost:    c.Host,
		HTTPClient: c.HTTPClient,
		Domain:     domain,
		RecordType: recordType,
		Name:       name,
	}
	if err := records.Get(conf); err != nil {
		return err
	}
	return nil
}

// PutRecord replaces all records for a given domain and (optionally) type and name
func (c *Client) PutRecord(domain, recordType, name, pointsto string, ttl int) error {
	return records.Put(&records.Config{
		APIKey:     c.apiKey,
		APISecret:  c.apiSecret,
		APIHost:    c.Host,
		HTTPClient: c.HTTPClient,
		Domain:     domain,
		RecordType: recordType,
		Name:       name,
		IP:         pointsto,
		TTL:        ttl,
	})
}
