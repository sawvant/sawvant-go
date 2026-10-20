package sawvant

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// StreamEvent represents an SSE event from the job stream.
type StreamEvent struct {
	Type string      // "progress", "completed", "failed", "error"
	Data JobResponse
}

// StreamJob streams job progress via Server-Sent Events.
//
// SSE events emitted by the server intentionally omit the full OptimizeResult
// to avoid "bufio.Scanner: token too long" errors on large result payloads.
// When a "completed" event is received, the SDK automatically fetches the full
// job via GET /v1/jobs/{id} and merges the result into the StreamEvent before
// delivering it to the caller. All other event types are forwarded as-is.
//
// The returned channels are both closed when the stream ends.
func (c *APIClient) StreamJob(ctx context.Context, jobID string) (<-chan StreamEvent, <-chan error) {
	events := make(chan StreamEvent)
	errs := make(chan error, 1)

	go func() {
		defer close(events)
		defer close(errs)

		basePath, err := c.cfg.ServerURL(0, nil)
		if err != nil {
			errs <- err
			return
		}
		url := fmt.Sprintf("%s/v1/jobs/%s/stream", basePath, jobID)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			errs <- err
			return
		}
		for key, value := range c.cfg.DefaultHeader {
			req.Header.Set(key, value)
		}
		req.Header.Set("Accept", "text/event-stream")

		resp, err := c.cfg.HTTPClient.Do(req)
		if err != nil {
			errs <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errs <- fmt.Errorf("stream failed: %d", resp.StatusCode)
			return
		}

		// SSE events are now small (no embedded result), so 64 KB is a
		// generous safety net. The 1 MB buffer that was here previously was a
		// workaround for the full result being embedded in the event data.
		scanner := bufio.NewScanner(resp.Body)
		scanner.Buffer(make([]byte, 0, bufio.MaxScanTokenSize), 1024*1024) // 1 MB max token size
		var currentEvent string

		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "event:") {
				currentEvent = strings.TrimSpace(line[6:])
			} else if strings.HasPrefix(line, "data:") {
				var data JobResponse
				if err := json.Unmarshal([]byte(strings.TrimSpace(line[5:])), &data); err != nil {
					errs <- fmt.Errorf("failed to parse event data: %w", err)
					return
				}

				// For "completed" events the server strips the Result field to
				// keep SSE payloads small. Fetch the full job via REST so the
				// caller always receives the complete JobResponse.
				if currentEvent == "completed" {
					full, _, fetchErr := c.JobsAPI.GetJobExecute(
						c.JobsAPI.GetJob(ctx, jobID),
					)
					if fetchErr != nil {
						errs <- fmt.Errorf("failed to fetch completed job result: %w", fetchErr)
						return
					}
					data = *full
				}

				select {
				case events <- StreamEvent{Type: currentEvent, Data: data}:
				case <-ctx.Done():
					return
				}
			}
		}

		if err := scanner.Err(); err != nil {
			errs <- err
		}
	}()

	return events, errs
}
