package main

// import client package
import (
	"github.com/zzstoatzz/prefect-go/pkg/client"
)

func main() {

	// create a new client
	prefect_client := client.New()

	// make a request with the client
	resp, err := prefect_client.Request("GET", "https://catfact.ninja/fact", nil)

	if err != nil {
		panic(err)
	}

	// print the response
	println(string(resp))

}
