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

	// List incident roles
	v, resp, err := client.IncidentRoles.List(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	for _, v := range v.IncidentRoles {
		fmt.Println(v.Id, v.Name)
	}

	fmt.Println("========================")

	// Get a single incident role
	v1, resp, err := client.IncidentRoles.Get(context.Background(), "<Role-ID>")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	fmt.Println(v1.IncidentRole.Name)
}
