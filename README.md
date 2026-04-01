# Sawvant Go SDK

The official Go client for the [Sawvant](https://sawvant.com) Cutting Optimization API.

- API version: 1.0.0
- Package: `sawvant`

## Installation

```sh
go get github.com/sawvant/sawvant-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	sawvant "github.com/sawvant/sawvant-go"
)

func main() {
	cfg := sawvant.NewConfiguration()
	cfg.Host = "api.sawvant.com"

	ctx := context.WithValue(context.Background(), sawvant.ContextAPIKeys, map[string]sawvant.APIKey{
		"ApiKeyAuth": {Key: "sk_your_api_key_here"},
	})

	client := sawvant.NewAPIClient(cfg)

	// Submit an optimization job
	req := sawvant.OptimizeRequest{
		Sheets: []sawvant.Sheet{
			{Width: 2440, Height: 1220, Quantity: 5},
		},
		Parts: []sawvant.Part{
			{Width: 600, Height: 400, Quantity: 10, Label: sawvant.PtrString("Panel A")},
			{Width: 300, Height: 200, Quantity: 20, Label: sawvant.PtrString("Panel B")},
		},
	}

	job, _, err := client.OptimizeApi.CreateOptimization(ctx).OptimizeRequest(req).Execute()
	if err != nil {
		log.Fatalf("create optimization: %v", err)
	}
	fmt.Printf("Job created: %s\n", job.GetId())

	// Poll for results
	for {
		result, _, err := client.JobsApi.GetJob(ctx, job.GetId()).Execute()
		if err != nil {
			log.Fatalf("get job: %v", err)
		}

		switch result.GetStatus() {
		case "completed":
			fmt.Printf("Optimization complete: %v\n", result.GetResult())
			return
		case "failed":
			log.Fatalf("job failed")
		default:
			time.Sleep(time.Second)
		}
	}
}
```

## SSE Streaming

Stream real-time progress events as the job runs:

```go
events, err := client.JobsApi.StreamJob(ctx, job.GetId()).Execute()
if err != nil {
	log.Fatalf("stream job: %v", err)
}

for event := range events {
	switch event.Type {
	case "progress":
		fmt.Printf("Progress: %v\n", event.Data)
	case "completed":
		fmt.Printf("Done: %v\n", event.Data)
		return
	case "failed":
		log.Fatalf("job failed: %v", event.Data)
	}
}
```

## Configuration

```go
cfg := sawvant.NewConfiguration()
cfg.Host = "api.sawvant.com"
cfg.Scheme = "https"

// Pass the API key via context on each request
ctx := context.WithValue(context.Background(), sawvant.ContextAPIKeys, map[string]sawvant.APIKey{
	"ApiKeyAuth": {Key: "sk_your_api_key_here"},
})
```

| Field | Description |
|-------|-------------|
| `cfg.Host` | API hostname (default: `api.sawvant.com`) |
| `cfg.Scheme` | URL scheme (default: `https`) |
| `cfg.HTTPClient` | Custom `*http.Client` |
| `ContextAPIKeys` | API key passed via context |

## API Reference

All endpoints are relative to `https://api.sawvant.com`.

| Method | HTTP | Path | Description |
|--------|------|------|-------------|
| `CreateOptimization` | POST | `/v1/optimize` | Submit a new cutting optimization job |
| `GetJob` | GET | `/v1/jobs/{id}` | Retrieve job status and result |
| `StreamJob` | GET | `/v1/jobs/{id}/stream` | Stream job progress via SSE |
| `GetHealth` | GET | `/health` | Health check (no auth required) |

## License

MIT
