package main

// import client package
import (
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

	// test health check
	resp, err := prefect_client.HealthCheck()

	// test hello
	// resp, err := prefect_client.Hello()

	// // test create flow
	// flowId := "103c9a44-cf7d-483e-bed2-681112b542ab"

	// resp, err := prefect_client.ReadFlow(flowId)

	if err != nil {
		println(err)
		panic(err)
	}

	// print the response
	println(resp)

}
