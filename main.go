package main

// import the base client from the client package
import "github.com/zzstoatzz/prefect-go/client/base"

func main() {
	PrefectClient := base.New()

	// make a request with the client
	resp, err := PrefectClient.Request("GET", "https://catfact.ninja/fact", nil)

	if err != nil {
		panic(err)
	}

	// print the response
	println(string(resp))

}
