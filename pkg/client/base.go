package client

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type PrefectClient struct {
	httpClient *http.Client
	BaseURL    string
	headers    map[string]string
}

// New creates a new PrefectClient
func NewClient(
	api_key string,
	baseURL string,
) *PrefectClient {

	headers := map[string]string{
		"Authorization": "Bearer " + api_key,
	}

	return &PrefectClient{
		httpClient: &http.Client{},
		BaseURL:    baseURL,
		headers:    headers,
	}
}

func (c *PrefectClient) Request(
	method string,
	url string,
	headers map[string]string,
	payload []byte,
) ([]byte, error) {

	req, _ := http.NewRequest(
		strings.ToUpper(method),
		url,
		strings.NewReader(string(payload)),
	)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	} else {
		return body, nil
	}
}
