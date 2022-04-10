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

	// List incidents
	v, resp, err := client.Incidents.List(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	for _, v := range v.Incidents {
		fmt.Println(v.Id, v.Name, v.PostmortemDocumentUrl)
	}

	fmt.Println("========================")

	// Get a single incident
	v1, resp, err := client.Incidents.Get(context.Background(), "<Incident-ID>")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	fmt.Println(v1.Incident.Name)
}
