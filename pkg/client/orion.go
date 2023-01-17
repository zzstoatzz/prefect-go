// define health check method for the client

// Path: pkg/client/orion.go

package client

import (
	"encoding/json"
	"errors"

	"github.com/zzstoatzz/prefect-go/pkg/data"
	"github.com/zzstoatzz/prefect-go/pkg/utils"
)

// admin methods

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

// read methods

func (client *PrefectClient) ReadFlow(flowId string) (*data.Flow, error) {

	if !utils.IsValidUUID(flowId) {
		return nil, errors.New("Provided `flowId` is not a valid UUID")
	}

	resp, err := client.Request(
		"GET", client.BaseURL+"/flows/"+flowId, client.headers, nil,
	)

	if err != nil {
		return nil, err
	}

	var flow data.Flow
	json.Unmarshal(resp, &flow)

	return &flow, nil
}

func (client *PrefectClient) ReadFlowByName(flowName string) (*data.Flow, error) {

	resp, err := client.Request(
		"GET", client.BaseURL+"/flows/name/"+flowName, client.headers, nil,
	)

	if err != nil {
		return nil, err
	}

	var flow data.Flow
	json.Unmarshal(resp, &flow)

	return &flow, nil
}

func (client *PrefectClient) ReadFlowRun(flowRunId string) (*data.FlowRun, error) {

	if !utils.IsValidUUID(flowRunId) {
		return nil, errors.New("Provided `flowRunId` is not a valid UUID")
	}

	resp, err := client.Request(
		"GET", client.BaseURL+"/flow_runs/"+flowRunId, client.headers, nil,
	)

	if err != nil {
		return nil, err
	}

	var flowRun data.FlowRun
	json.Unmarshal(resp, &flowRun)

	return &flowRun, nil
}
