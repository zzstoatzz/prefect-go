package main

// import client package
import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/zzstoatzz/prefect-go/pkg/client"
)

func main() {
	// load .env file
	godotenv.Load()

	api_key := os.Getenv("PREFECT_API_KEY")
	base_url := os.Getenv("PREFECT_API_URL")

	// create a new client
	prefect_client := client.NewClient(
		api_key,
		base_url,
	)

	// // test health check
	// resp, err := prefect_client.HealthCheck()

	// test hello
	// resp, err := prefect_client.Hello()

	// // test read flow
	// flowId := "103c9a44-cf7d-483e-bed2-681112b542ab"

	// resp, err := prefect_client.ReadFlow(flowId)

	// // test read flow by name
	// flowName := "testing"

	// resp, err := prefect_client.ReadFlowByName(flowName)

	// test read flow runs
	flowRunId := "c451b292-148e-4f47-8d8c-393712975936"

	resp, err := prefect_client.ReadFlowRun(flowRunId)

	// handle errors
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	// pretty print the response
	empJson, _ := json.MarshalIndent(resp, "", " ")

	fmt.Printf("RESPONSE: %s\n", string(empJson))

}
