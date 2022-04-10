package main

import (
	"context"
	"fmt"
	"os"

	"github.com/andygrunwald/go-incident"
)

func main() {
	apiKey := os.Getenv("INCIDENT_IO_API_KEY")
	client := incident.NewClient(apiKey, nil)

	// List severities
	v, resp, err := client.Severities.List(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	for _, v := range v.Severities {
		fmt.Println(v.Id, v.Name, v.Rank)
	}

	fmt.Println("========================")

	// Get a single severity
	v1, resp, err := client.Severities.Get(context.Background(), "<Severity-ID>")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	fmt.Println(v1.Severity.Name)
}
