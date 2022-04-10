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

	// Incidents with opts
	opt := &incident.IncidentsListOptions{
		PageSize: 5,
		Status: []string{
			incident.IncidentStatusClosed,
		},
	}

	// get all pages of results
	var allIncidents []incident.Incident
	for {
		fmt.Printf("Requesting incidents - After ID: %s\n", opt.After)
		incidents, _, err := client.Incidents.ListIncidents(context.Background(), opt)
		if err != nil {
			panic(err)
		}

		allIncidents = append(allIncidents, incidents.Incidents...)
		fmt.Printf("Appended %d items, in total now %d\n", len(incidents.Incidents), len(allIncidents))

		if incidents.PaginationMeta.TotalRecordCount == int64(len(allIncidents)) {
			break
		}
		opt.After = incidents.Incidents[len(incidents.Incidents)-1].Id
	}

	for _, v := range allIncidents {
		fmt.Println(v.Id, v.Name, v.PostmortemDocumentUrl)
	}
}
