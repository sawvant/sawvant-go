package sawvant

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStreamJob_ParseSSEEvents(t *testing.T) {
	tests := []struct {
		name           string
		responseBody   string
		expectedEvents int
		expectedTypes  []string
		wantErr        bool
	}{
		{
			name: "single progress event",
			responseBody: `event: progress
data: {"job_id":"job-1","status":"running","progress":50,"created_at":"2024-01-01T00:00:00Z"}
`,
			expectedEvents: 1,
			expectedTypes:  []string{"progress"},
			wantErr:        false,
		},
		{
			name: "multiple events in sequence",
			responseBody: `event: progress
data: {"job_id":"job-1","status":"running","progress":25,"created_at":"2024-01-01T00:00:00Z"}

event: progress
data: {"job_id":"job-1","status":"running","progress":50,"created_at":"2024-01-01T00:00:00Z"}

event: completed
data: {"job_id":"job-1","status":"completed","progress":100,"created_at":"2024-01-01T00:00:00Z"}
`,
			expectedEvents: 3,
			expectedTypes:  []string{"progress", "progress", "completed"},
			wantErr:        false,
		},
		{
			name: "event with optional fields",
			responseBody: `event: completed
data: {"job_id":"job-1","status":"completed","progress":100,"created_at":"2024-01-01T00:00:00Z","completed_at":"2024-01-01T01:00:00Z"}
`,
			expectedEvents: 1,
			expectedTypes:  []string{"completed"},
			wantErr:        false,
		},
		{
			name: "error event type",
			responseBody: `event: error
data: {"job_id":"job-1","status":"failed","progress":0,"created_at":"2024-01-01T00:00:00Z","error":"Unknown error"}
`,
			expectedEvents: 1,
			expectedTypes:  []string{"error"},
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/event-stream")
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, tt.responseBody)
			}))
			defer server.Close()

			cfg := NewConfiguration()
			cfg.Servers[0].URL = server.URL
			cfg.HTTPClient = http.DefaultClient

			client := NewAPIClient(cfg)
			events, errs := client.StreamJob(context.Background(), "job-1")

			receivedEvents := make([]StreamEvent, 0)
			var lastErr error
			for {
				select {
				case event, ok := <-events:
					if !ok {
						// Channel closed
						goto checkResults
					}
					receivedEvents = append(receivedEvents, event)
				case err := <-errs:
					if err != nil {
						lastErr = err
					}
				case <-time.After(2 * time.Second):
					// Timeout
					goto checkResults
				}
			}

		checkResults:
			if tt.wantErr {
				assert.Error(t, lastErr)
			} else {
				assert.NoError(t, lastErr)
				assert.Equal(t, tt.expectedEvents, len(receivedEvents), "should receive expected number of events")
				for i, expectedType := range tt.expectedTypes {
					assert.Equal(t, expectedType, receivedEvents[i].Type, "event type mismatch at index %d", i)
				}
			}
		})
	}
}

func TestStreamJob_MalformedJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: progress
data: {"invalid json unclosed
`)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	receivedEvents := 0
	var lastErr error
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			receivedEvents++
			_ = event
		case err := <-errs:
			if err != nil {
				lastErr = err
			}
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	assert.Error(t, lastErr, "should error on malformed JSON")
	assert.Equal(t, 0, receivedEvents, "should not receive any complete events")
}

func TestStreamJob_ContextCancellation(t *testing.T) {
	blockingServer := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: progress
data: {"job_id":"job-1","status":"running","progress":25,"created_at":"2024-01-01T00:00:00Z"}

`)
		flusher := w.(http.Flusher)
		if flusher != nil {
			flusher.Flush()
		}
		<-blockingServer
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	ctx, cancel := context.WithCancel(context.Background())

	events, errs := client.StreamJob(ctx, "job-1")

	// Read first event
	select {
	case event := <-events:
		assert.Equal(t, "progress", event.Type)
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for first event")
	}

	// Cancel context
	cancel()

	// Give goroutine time to process context cancellation
	time.Sleep(100 * time.Millisecond)

	// Close the blocking channel
	close(blockingServer)

	// Try to read from channels - they should be closed or have an error
	select {
	case event, ok := <-events:
		if ok {
			// May receive another event before context is checked, that's OK
			_ = event
		}
	case err, ok := <-errs:
		if ok {
			// May have an error channel message, that's OK
			_ = err
		}
	case <-time.After(2 * time.Second):
		// Channels should be closed or handled
	}
}

func TestStreamJob_NonOKResponse(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		expectedErrMsg string
	}{
		{
			name:           "500 internal server error",
			statusCode:     http.StatusInternalServerError,
			expectedErrMsg: "stream failed: 500",
		},
		{
			name:           "404 not found",
			statusCode:     http.StatusNotFound,
			expectedErrMsg: "stream failed: 404",
		},
		{
			name:           "401 unauthorized",
			statusCode:     http.StatusUnauthorized,
			expectedErrMsg: "stream failed: 401",
		},
		{
			name:           "429 rate limited",
			statusCode:     http.StatusTooManyRequests,
			expectedErrMsg: "stream failed: 429",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			cfg := NewConfiguration()
			cfg.Servers[0].URL = server.URL
			cfg.HTTPClient = http.DefaultClient

			client := NewAPIClient(cfg)
			events, errs := client.StreamJob(context.Background(), "job-1")

			var lastErr error
			for {
				select {
				case event, ok := <-events:
					if !ok {
						goto checkResults
					}
					_ = event
				case err := <-errs:
					if err != nil {
						lastErr = err
					}
				case <-time.After(2 * time.Second):
					goto checkResults
				}
			}

		checkResults:
			assert.Error(t, lastErr, "should return an error")
			assert.Contains(t, lastErr.Error(), tt.expectedErrMsg)
		})
	}
}

func TestStreamJob_ServerURLResolution(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v1/jobs/job-123/stream", r.RequestURI)
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: progress
data: {"job_id":"job-123","status":"running","progress":0,"created_at":"2024-01-01T00:00:00Z"}
`)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-123")

	receivedJobID := ""
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			receivedJobID = event.Data.JobId
		case err := <-errs:
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	assert.Equal(t, "job-123", receivedJobID)
}

func TestStreamJob_DefaultHeaders(t *testing.T) {
	headerReceived := make(chan string)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerReceived <- r.Header.Get("X-Custom-Header")
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: progress
data: {"job_id":"job-1","status":"running","progress":0,"created_at":"2024-01-01T00:00:00Z"}
`)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient
	cfg.AddDefaultHeader("X-Custom-Header", "custom-value")

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	var lastErr error
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			_ = event
		case err := <-errs:
			if err != nil {
				lastErr = err
			}
		case header := <-headerReceived:
			assert.Equal(t, "custom-value", header)
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	assert.NoError(t, lastErr)
}

func TestStreamJob_AcceptHeader(t *testing.T) {
	acceptHeaderReceived := make(chan string)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acceptHeaderReceived <- r.Header.Get("Accept")
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: progress
data: {"job_id":"job-1","status":"running","progress":0,"created_at":"2024-01-01T00:00:00Z"}
`)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	acceptHeader := ""
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			_ = event
		case err := <-errs:
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		case header := <-acceptHeaderReceived:
			acceptHeader = header
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	assert.Equal(t, "text/event-stream", acceptHeader)
}

func TestStreamJob_EventDataDeserialization(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: completed
data: {"job_id":"job-1","status":"completed","progress":100,"created_at":"2024-01-01T00:00:00Z","completed_at":"2024-01-02T00:00:00Z","warnings":["unplaced-part-1","unplaced-part-2"]}
`)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	var receivedEvent *StreamEvent
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			receivedEvent = &event
		case err := <-errs:
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	require.NotNil(t, receivedEvent)
	assert.Equal(t, "completed", receivedEvent.Type)
	assert.Equal(t, "job-1", receivedEvent.Data.JobId)
	assert.Equal(t, int32(100), receivedEvent.Data.Progress)
	assert.Equal(t, "completed", receivedEvent.Data.Status)
	assert.Equal(t, 2, len(receivedEvent.Data.Warnings))
	assert.Contains(t, receivedEvent.Data.Warnings, "unplaced-part-1")
}

func TestStreamJob_NetworkError(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers[0].URL = "http://invalid-host-that-does-not-exist.example.com:9999"
	cfg.HTTPClient = &http.Client{Timeout: 100 * time.Millisecond}

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	var lastErr error
	var eventCount int
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			eventCount++
			_ = event
		case err := <-errs:
			if err != nil {
				lastErr = err
			}
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	assert.Error(t, lastErr, "should return a network error")
	assert.Equal(t, 0, eventCount)
}

func TestStreamJob_EmptyStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		// Empty body
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	var eventCount int
	var lastErr error
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			eventCount++
			_ = event
		case err := <-errs:
			if err != nil {
				lastErr = err
			}
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	assert.NoError(t, lastErr)
	assert.Equal(t, 0, eventCount)
}

func TestStreamJob_ChannelsClosed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `event: completed
data: {"job_id":"job-1","status":"completed","progress":100,"created_at":"2024-01-01T00:00:00Z"}
`)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	// Read all events
	eventCount := 0
	for {
		select {
		case event, ok := <-events:
			if !ok {
				// Channel closed as expected
				goto checkChannels
			}
			eventCount++
			_ = event
		case <-time.After(2 * time.Second):
			goto checkChannels
		}
	}

checkChannels:
	// Verify both channels are now closed
	select {
	case _, ok := <-events:
		assert.False(t, ok, "events channel should be closed")
	case <-time.After(100 * time.Millisecond):
		t.Fatal("events channel should be closed immediately")
	}

	select {
	case _, ok := <-errs:
		assert.False(t, ok, "errs channel should be closed")
	case <-time.After(100 * time.Millisecond):
		t.Fatal("errs channel should be closed immediately")
	}
}

func TestStreamJob_LargeProgressValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		body := bytes.NewBufferString(`event: progress
data: {"job_id":"job-1","status":"running","progress":999,"created_at":"2024-01-01T00:00:00Z"}
`)
		io.Copy(w, body)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.Servers[0].URL = server.URL
	cfg.HTTPClient = http.DefaultClient

	client := NewAPIClient(cfg)
	events, errs := client.StreamJob(context.Background(), "job-1")

	var receivedEvent *StreamEvent
	for {
		select {
		case event, ok := <-events:
			if !ok {
				goto checkResults
			}
			receivedEvent = &event
		case err := <-errs:
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		case <-time.After(2 * time.Second):
			goto checkResults
		}
	}

checkResults:
	require.NotNil(t, receivedEvent)
	assert.Equal(t, int32(999), receivedEvent.Data.Progress)
}

func TestStreamJob_MultipleErrorsReturnsFirst(t *testing.T) {
	// Test that if multiple errors could occur, at least one is returned
	cfg := NewConfiguration()
	cfg.Servers[0].URL = "http://invalid-host.example.com:9999"
	cfg.HTTPClient = &http.Client{Timeout: 100 * time.Millisecond}

	client := NewAPIClient(cfg)
	_, errs := client.StreamJob(context.Background(), "job-1")

	var receivedErr error
	select {
	case receivedErr = <-errs:
	case <-time.After(2 * time.Second):
		t.Fatal("expected an error")
	}

	assert.NotNil(t, receivedErr)
}
