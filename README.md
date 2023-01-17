# a golang prefect client

## Overview
Very much a work in progress. Currently only supports the following endpoints:
- health check: `/health` (GET)

- read flow by ID: `/flows/{flow_id}` (GET)

- read flow by name: `/flows/name/{flow_name}` (GET)

- read flow run by ID: `/flow-runs/{flow_run_id}` (GET)

```go 
package main

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
	prefect_client := client.NewClient(api_key, base_url)

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
```

I would like to generate the contents of the [`pkg/data`](pkg/data/) directory using the [OpenAPI spec](https://api.prefect.cloud/api/openapi.json). I have not yet figured out how to do this, although something in [this](https://www.jvt.me/posts/2022/04/06/generate-go-struct-openapi/) vein seems promising.


## Development
### Setup
```bash
git clone https://github.com/zzstoatzz/prefect-go.git
```