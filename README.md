# go-incident: Go client library for [Incident.io](https://incident.io/)

[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/andygrunwald/go-incident)
[![Test Status](https://github.com/google/go-github/workflows/tests/badge.svg)](https://github.com/andygrunwald/go-incident/actions?query=workflow%3Atesting)

Go client library for accessing the [Incident.io](https://incident.io/) [API](https://api-docs.incident.io/).

## Installation

go-incident is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/andygrunwald/go-incident
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/andygrunwald/go-incident"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/andygrunwald/go-incident@main
```

## Usage

```go
import "github.com/andygrunwald/go-incident"
```

Construct a new Incident.io client, then use the various services on the client to access different parts of the [Incident.io API](https://api-docs.incident.io/).
For example:

```go
apiKey := "<my-secret-api-key>"
client := incident.NewClient(apiKey, nil)

// List all incidents for your organisation.
incidents, response, err := c.Incidents.ListIncidents(context.Background(), nil)
```

Some API methods have optional parameters that can be passed. For example:

```go
apiKey := "<my-secret-api-key>"
client := incident.NewClient(apiKey, nil)

// List only closed incidents for your organisation in page chunks of 5.
opt := &incident.IncidentsListOptions{
    PageSize: 5,
    Status: []string{
        incident.IncidentStatusClosed,
    },
}
incidents, response, err := c.Incidents.ListIncidents(context.Background(), opt)
```

The services of a client divide the API into logical chunks and correspond to the structure of the [Incident.io API](https://api-docs.incident.io/) documentation .

NOTE: Using the [context](https://pkg.go.dev/context) package, one can easily pass cancelation signals and deadlines to various services of the client for handling a request.
In case there is no context available, then `context.Background()` can be used as a starting point.

For more sample code snippets, head over to the [example](https://github.com/google/go-github/tree/master/example) directory.

### Authentication

For all requests made to the incident.io API, you'll need an API key.
Right now, there is no public incident.io API.

To create an API key, head to the incident dashboard and visit [API keys](https://app.incident.io/settings/api-keys).

The API key will be passed as the first argument, when constructing a new client:

```go
apiKey := "<my-secret-api-key>"
client := incident.NewClient(apiKey, nil)
```

### Errors

Errors provided by the Incident.io API will be mapped to the [ErrorResponse](https://pkg.go.dev/github.com/andygrunwald/go-incident#ErrorResponse) type and can be investigated further:

```go
// Do a API call ...
if err != nil {
    if responseErr, ok := err.(*incident.ErrorResponse); ok {
        // Do something with responseErr, like printing
        fmt.Printf("%+v", responseErr.Type)
    }
}
```

All error details provided by the API are available.
See [Making requests > Errors in the Incident.ip API docs](https://api-docs.incident.io/#section/Making-requests/Errors) for more details.

### Pagination

Some requests support pagination.
Pagination options are described in the options per API call once supported.
The returned data contains a [PaginationMeta](https://pkg.go.dev/github.com/andygrunwald/go-incident#PaginationMeta) struct with paging information.

```go
apiKey := "<my-secret-api-key>"
client := incident.NewClient(apiKey, nil)

opt := &incident.IncidentsListOptions{
    PageSize: 5,
}

// Get all pages of incidents
var allIncidents []incident.Incident
for {
    incidents, _, err := client.Incidents.ListIncidents(context.Background(), opt)
    if err != nil {
        panic(err)
    }
    allIncidents = append(allIncidents, incidents.Incidents...)

    // Calculate if there is a next page
    if incidents.PaginationMeta.TotalRecordCount == int64(len(allIncidents)) {
        break
    }
    opt.After = incidents.Incidents[len(incidents.Incidents)-1].Id
}
```

## Contributing

I would like to cover the entire Incident.io API and contributions are of course always welcome.
The calling pattern is pretty well established, so adding new methods is relatively straightforward.

## Inspired by

The structure, code and documentation of this project is inspired by [google/go-github](https://github.com/google/go-github).

## License

This library is distributed under the MIT License found in the [LICENSE](./LICENSE) file.