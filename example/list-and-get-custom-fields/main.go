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

	// List custom fields
	v, resp, err := client.CustomFields.ListCustomFields(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	for _, v := range v.CustomFields {
		fmt.Println(v.Id, v.Name, v.FieldType)
	}

	fmt.Println("========================")

	// Get a single custom field
	v1, resp, err := client.CustomFields.GetCustomField(context.Background(), "<Custom-Field-ID>")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response: %v\n", resp.Status)

	fmt.Println(v1.CustomField.Name)
}
