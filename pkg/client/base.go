package client

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type PrefectClient struct {
	httpClient *http.Client
}

func New() *PrefectClient {
	return &PrefectClient{
		httpClient: &http.Client{},
	}
}

func (c *PrefectClient) Request(
	method string,
	url string,
	headers map[string]string,
) ([]byte, error) {
	req, _ := http.NewRequest(strings.ToUpper(method), url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
