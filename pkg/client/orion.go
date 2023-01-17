// define health check method for the client

// Path: pkg/client/orion.go

package client

import (
	"errors"
)

func (client *PrefectClient) HealthCheck() (bool, error) {

	resp, err := client.Request("GET", client.BaseURL+"/health", client.headers, nil)

	if string(resp) != "true" || err != nil {

		return false, errors.New("health check failed")

	}

	return true, nil
}

func (client *PrefectClient) Hello() (string, error) {

	resp, err := client.Request("GET", client.BaseURL+"/hello", client.headers, nil)

	if err != nil {

		return "", err

	}

	return string(resp), nil
}

func (client *PrefectClient) ReadFlow(flowId string) (string, error) {

	resp, err := client.Request(
		"GET", client.BaseURL+"/flows/"+flowId, client.headers, nil,
	)

	if err != nil {
		println(err)
		return "", err

	}

	return string(resp), nil
}
