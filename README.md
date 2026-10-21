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
	client := sawvant.NewAPIClient(sawvant.NewConfiguration())

	ctx := context.WithValue(context.Background(), sawvant.ContextAPIKeys, map[string]sawvant.APIKey{
		"apiKey": {Key: "sk_your_api_key_here"},
	})

	// Submit an optimization job. Sizes are in mm.
	req := sawvant.OptimizeRequest{
		Sheets: []sawvant.Sheet{
			{Id: "board", Length: 2800, Width: 2070, Quantity: 10, Grain: sawvant.NONE},
		},
		Parts: []sawvant.Part{
			{Id: "panel-a", Length: 600, Width: 400, Quantity: 10, Grain: sawvant.NONE},
			{Id: "panel-b", Length: 300, Width: 200, Quantity: 20, Grain: sawvant.NONE},
		},
		Machine: sawvant.Machine{BladeThickness: 4.4, MaxLevels: 3, CutDirection: sawvant.DEFAULT},
	}

	job, _, err := client.OptimizationAPI.CreateOptimization(ctx).OptimizeRequest(req).Execute()
	if err != nil {
		log.Fatalf("create optimization: %v", err)
	}
	fmt.Printf("Job created: %s\n", job.GetJobId())

	// Poll for results
	for {
		result, _, err := client.JobsAPI.GetJob(ctx, job.GetJobId()).Execute()
		if err != nil {
			log.Fatalf("get job: %v", err)
		}

		switch result.GetStatus() {
		case "completed":
			fmt.Printf("Optimization complete: %+v\n", result.GetResult())
			return
		case "failed":
			log.Fatalf("job failed: %s", result.GetError())
		default:
			time.Sleep(time.Second)
		}
	}
}
```

## SSE Streaming

Stream real-time progress events as the job runs with `client.StreamJob`. It returns an events channel and an errors channel; both close when the stream ends.

The stream request sends the headers from `cfg.DefaultHeader`, not the API key from the context, so set the key as a default header. Pass the job ID that `CreateOptimization` returned:

```go
package main

import (
	"context"
	"fmt"
	"log"

	sawvant "github.com/sawvant/sawvant-go"
)

func main() {
	cfg := sawvant.NewConfiguration()
	cfg.AddDefaultHeader("X-API-Key", "sk_your_api_key_here")
	client := sawvant.NewAPIClient(cfg)

	jobID := "550e8400-e29b-41d4-a716-446655440000"
	events, errs := client.StreamJob(context.Background(), jobID)

	for event := range events {
		switch event.Type {
		case "progress":
			fmt.Printf("Progress: %d%%\n", event.Data.GetProgress())
		case "completed":
			fmt.Printf("Done: %+v\n", event.Data.GetResult())
		case "failed":
			log.Fatalf("job failed: %s", event.Data.GetError())
		}
	}

	if err := <-errs; err != nil {
		log.Fatalf("stream job: %v", err)
	}
}
```

Cancel the context to stop the stream early.

When a `completed` event arrives, the SDK fetches the full job via `GET /v1/jobs/{id}` and puts it in the event, so `event.Data.GetResult()` holds the complete result. The server leaves the result out of stream events to keep them small. Other events are passed on as received.

## Configuration

`sawvant.NewConfiguration()` points at `https://api.sawvant.com`. Pass the API key per request through the context, as in the Quick Start, under the key `apiKey`.

| Field | Description |
|-------|-------------|
| `cfg.Host` | API hostname (default: `api.sawvant.com`) |
| `cfg.Scheme` | URL scheme (default: `https`) |
| `cfg.HTTPClient` | Custom `*http.Client` |
| `cfg.DefaultHeader` | Headers sent with every request, set with `cfg.AddDefaultHeader` |
| `sawvant.ContextAPIKeys` | Context key for the API key map |

## API Reference

All endpoints are relative to `https://api.sawvant.com`.

| Service | Method | HTTP | Path | Description |
|---------|--------|------|------|-------------|
| `OptimizationAPI` | `CreateOptimization` | POST | `/v1/optimize` | Submit a new cutting optimization job |
| `OptimizationAPI` | `CreateBatchOptimization` | POST | `/v1/optimize/batch` | Submit several optimization jobs at once |
| `JobsAPI` | `GetJob` | GET | `/v1/jobs/{id}` | Retrieve job status and result |
| `JobsAPI` | `GetJobDiagramSVG` | GET | `/v1/jobs/{id}/diagram` | Cutting diagram as SVG |
| `JobsAPI` | `GetJobDiagramPDF` | GET | `/v1/jobs/{id}/diagram.pdf` | Cutting diagram as PDF |
| `client` | `StreamJob` | GET | `/v1/jobs/{id}/stream` | Stream job progress via SSE |
| `SystemAPI` | `GetHealth` | GET | `/health` | Health check (no auth required) |
| `SystemAPI` | `ListTiers` | GET | `/v1/tiers` | List subscription tiers (no auth required) |

## License

Apache-2.0. See [LICENSE](LICENSE).
