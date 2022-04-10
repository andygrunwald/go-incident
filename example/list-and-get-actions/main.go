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

	// List actions
	v, resp, err := client.Actions.ListActions(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	for _, v := range v.Actions {
		fmt.Println(v.Id, v.Description, v.Status)
	}

	fmt.Println("========================")

	// Get a single action
	v1, resp, err := client.Actions.GetAction(context.Background(), "<Actions-ID>")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	fmt.Println(v1.Action.Description)
}
