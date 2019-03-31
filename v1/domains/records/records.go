package records

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	APIKey     string
	APISecret  string
	APIHost    string
	HTTPClient *http.Client
	Domain     string
	IP         string
	RecordType string
	Name       string
	TTL        int
}

type record struct {
	Data     string `json:"data"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Protocol string `json:"protocol"`
	Service  string `json:"service"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
	Weight   int    `json:"weight"`
}

// Get : GET /v1/domains/{domain}/records/{type}/{name}
func Get(c *Config) error {
	req, err := buildRequest(c.APIHost, c.Domain, c.RecordType, c.Name, http.MethodGet, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf(" sso-key %s:%s", c.APIKey, c.APISecret))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	fmt.Println(string(respBytes))
	return nil
}

// Put : PUT /v1/domains/{domain}/records/{type}/{name}
func Put(c *Config) error {
	req, err := buildRequest(c.APIHost, c.Domain, c.RecordType, c.Name, http.MethodPut, &record{
		Data:     c.IP,
		Name:     c.Name,
		TTL:      600,
		Type:     c.RecordType,
		Weight:   1,
		Port:     1,
		Priority: 0,
	})
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf(" sso-key %s:%s", c.APIKey, c.APISecret))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	fmt.Println(string(respBytes))
	return nil
}

func buildRequest(host, domain, recordType, name, method string, rec *record) (*http.Request, error) {
	url := fmt.Sprintf("%s/v1/domains/%s/records", host, domain)
	if recordType != "" {
		url = fmt.Sprintf("%s/%s", url, recordType)
	}
	if name != "" {
		url = fmt.Sprintf("%s/%s", url, name)
	}
	var req *http.Request
	var err error
	if method == http.MethodPut {
		byt, err := json.Marshal([]record{
			*rec,
		})
		if err != nil {
			return nil, err
		}
		if req, err = http.NewRequest(method, url, bytes.NewBuffer(byt)); err != nil {
			return nil, err
		}
	} else {
		if req, err = http.NewRequest(method, url, nil); err != nil {
			return nil, err
		}
	}

	return req, nil
}
