package records

import (
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
	RecordType string
	Name       string
}

// Get : GET /v1/domains/{domain}/records/{type}/{name}
func Get(c *Config) error {
	req, err := buildGetRecordsRequest(c.APIHost, c.Domain, c.RecordType, c.Name)
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

func buildGetRecordsRequest(host, domain, recordType, name string) (*http.Request, error) {
	url := fmt.Sprintf("%s/v1/domains/%s/records", host, domain)
	if recordType != "" {
		url = fmt.Sprintf("%s/%s", url, recordType)
	}
	if name != "" {
		url = fmt.Sprintf("%s/%s", url, name)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
