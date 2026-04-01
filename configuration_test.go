package sawvant

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfiguration_Defaults(t *testing.T) {
	cfg := NewConfiguration()

	assert.NotNil(t, cfg)
	assert.Equal(t, "sawvant-go-sdk/1.0.0", cfg.UserAgent)
	assert.False(t, cfg.Debug)
	assert.Equal(t, "", cfg.Host)
	assert.Equal(t, "", cfg.Scheme)
	assert.NotNil(t, cfg.DefaultHeader)
	assert.Equal(t, 0, len(cfg.DefaultHeader))
	assert.NotNil(t, cfg.Servers)
	assert.Greater(t, len(cfg.Servers), 0)
	assert.Equal(t, "https://api.sawvant.com", cfg.Servers[0].URL)
	assert.Equal(t, "Production", cfg.Servers[0].Description)
}

func TestNewConfiguration_HTTPClient(t *testing.T) {
	cfg := NewConfiguration()

	// HTTP client should be nil initially
	assert.Nil(t, cfg.HTTPClient)

	// Set a custom HTTP client
	customClient := &http.Client{}
	cfg.HTTPClient = customClient
	assert.Equal(t, customClient, cfg.HTTPClient)
}

func TestAddDefaultHeader(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		value  string
		verify func(t *testing.T, cfg *Configuration)
	}{
		{
			name:  "add single header",
			key:   "X-Custom-Header",
			value: "custom-value",
			verify: func(t *testing.T, cfg *Configuration) {
				assert.Equal(t, "custom-value", cfg.DefaultHeader["X-Custom-Header"])
			},
		},
		{
			name:  "add authorization header",
			key:   "Authorization",
			value: "Bearer token123",
			verify: func(t *testing.T, cfg *Configuration) {
				assert.Equal(t, "Bearer token123", cfg.DefaultHeader["Authorization"])
			},
		},
		{
			name:  "overwrite existing header",
			key:   "X-API-Key",
			value: "new-key",
			verify: func(t *testing.T, cfg *Configuration) {
				assert.Equal(t, "new-key", cfg.DefaultHeader["X-API-Key"])
				// Verify by adding it twice
				cfg.AddDefaultHeader("X-API-Key", "old-key")
				assert.Equal(t, "old-key", cfg.DefaultHeader["X-API-Key"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := NewConfiguration()
			cfg.AddDefaultHeader(tt.key, tt.value)
			tt.verify(t, cfg)
		})
	}
}

func TestServerURL(t *testing.T) {
	tests := []struct {
		name      string
		index     int
		variables map[string]string
		wantErr   bool
		expected  string
	}{
		{
			name:      "default server at index 0",
			index:     0,
			variables: nil,
			wantErr:   false,
			expected:  "https://api.sawvant.com",
		},
		{
			name:      "invalid index negative",
			index:     -1,
			variables: nil,
			wantErr:   true,
		},
		{
			name:      "invalid index out of range",
			index:     10,
			variables: nil,
			wantErr:   true,
		},
		{
			name:      "server at index 0 with nil variables",
			index:     0,
			variables: nil,
			wantErr:   false,
			expected:  "https://api.sawvant.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := NewConfiguration()
			url, err := cfg.ServerURL(tt.index, tt.variables)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, url)
			}
		})
	}
}

func TestServerURL_MultipleServers(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
		{
			URL:         "https://staging.sawvant.com",
			Description: "Staging",
		},
		{
			URL:         "http://localhost:8080",
			Description: "Local",
		},
	}

	tests := []struct {
		name     string
		index    int
		expected string
		wantErr  bool
	}{
		{
			name:     "first server",
			index:    0,
			expected: "https://api.sawvant.com",
			wantErr:  false,
		},
		{
			name:     "second server",
			index:    1,
			expected: "https://staging.sawvant.com",
			wantErr:  false,
		},
		{
			name:     "third server",
			index:    2,
			expected: "http://localhost:8080",
			wantErr:  false,
		},
		{
			name:     "invalid index",
			index:    3,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := cfg.ServerURL(tt.index, nil)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, url)
			}
		})
	}
}

func TestServerURL_WithVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://{host}.sawvant.com",
			Description: "Dynamic host",
			Variables: map[string]ServerVariable{
				"host": {
					Description:  "Hostname",
					DefaultValue: "api",
					EnumValues:   []string{"api", "staging", "dev"},
				},
			},
		},
	}

	tests := []struct {
		name      string
		variables map[string]string
		expected  string
		wantErr   bool
	}{
		{
			name:      "use default variable",
			variables: nil,
			expected:  "https://api.sawvant.com",
			wantErr:   false,
		},
		{
			name:      "override with valid variable",
			variables: map[string]string{"host": "staging"},
			expected:  "https://staging.sawvant.com",
			wantErr:   false,
		},
		{
			name:      "override with dev variable",
			variables: map[string]string{"host": "dev"},
			expected:  "https://dev.sawvant.com",
			wantErr:   false,
		},
		{
			name:      "invalid enum value",
			variables: map[string]string{"host": "invalid"},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url, err := cfg.ServerURL(0, tt.variables)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, url)
			}
		})
	}
}

func TestServerURLWithContext_NoContext(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
	}

	url, err := cfg.ServerURLWithContext(nil, "SomeEndpoint")

	assert.NoError(t, err)
	assert.Equal(t, "https://api.sawvant.com", url)
}

func TestServerURLWithContext_WithServerIndex(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
		{
			URL:         "https://staging.sawvant.com",
			Description: "Staging",
		},
	}

	ctx := context.WithValue(context.Background(), ContextServerIndex, 1)
	url, err := cfg.ServerURLWithContext(ctx, "SomeEndpoint")

	assert.NoError(t, err)
	assert.Equal(t, "https://staging.sawvant.com", url)
}

func TestServerURLWithContext_WithOperationServerIndex(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
		{
			URL:         "https://staging.sawvant.com",
			Description: "Staging",
		},
	}
	cfg.OperationServers = map[string]ServerConfigurations{
		"JobsAPIService.GetJob": ServerConfigurations{
			{
				URL:         "http://localhost:8080",
				Description: "Local",
			},
		},
	}

	operationIndices := map[string]int{
		"JobsAPIService.GetJob": 0,
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerIndices, operationIndices)
	url, err := cfg.ServerURLWithContext(ctx, "JobsAPIService.GetJob")

	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:8080", url)
}

func TestServerURLWithContext_WithServerVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://{host}.sawvant.com",
			Description: "Dynamic",
			Variables: map[string]ServerVariable{
				"host": {
					Description:  "Hostname",
					DefaultValue: "api",
					EnumValues:   []string{"api", "staging"},
				},
			},
		},
	}

	variables := map[string]string{"host": "staging"}
	ctx := context.WithValue(context.Background(), ContextServerVariables, variables)
	url, err := cfg.ServerURLWithContext(ctx, "SomeEndpoint")

	assert.NoError(t, err)
	assert.Equal(t, "https://staging.sawvant.com", url)
}

func TestServerURLWithContext_WithOperationServerVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
	}
	cfg.OperationServers = map[string]ServerConfigurations{
		"JobsAPIService.StreamJob": ServerConfigurations{
			{
				URL:         "https://{env}.sawvant.com",
				Description: "Stream endpoint",
				Variables: map[string]ServerVariable{
					"env": {
						Description:  "Environment",
						DefaultValue: "api",
						EnumValues:   []string{"api", "stream"},
					},
				},
			},
		},
	}

	opVars := map[string]map[string]string{
		"JobsAPIService.StreamJob": {
			"env": "stream",
		},
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerVariables, opVars)
	url, err := cfg.ServerURLWithContext(ctx, "JobsAPIService.StreamJob")

	assert.NoError(t, err)
	assert.Equal(t, "https://stream.sawvant.com", url)
}

func TestServerURLWithContext_InvalidContextValue(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
	}

	// Pass invalid type for ContextServerIndex
	ctx := context.WithValue(context.Background(), ContextServerIndex, "invalid")
	_, err := cfg.ServerURLWithContext(ctx, "SomeEndpoint")

	assert.Error(t, err)
}

func TestServerURLWithContext_InvalidOperationServerIndices(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
	}

	// Pass invalid type for ContextOperationServerIndices
	ctx := context.WithValue(context.Background(), ContextOperationServerIndices, "invalid")
	_, err := cfg.ServerURLWithContext(ctx, "SomeEndpoint")

	assert.Error(t, err)
}

func TestServerURLWithContext_InvalidServerVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
	}

	// Pass invalid type for ContextServerVariables
	ctx := context.WithValue(context.Background(), ContextServerVariables, "invalid")
	_, err := cfg.ServerURLWithContext(ctx, "SomeEndpoint")

	assert.Error(t, err)
}

func TestServerURLWithContext_InvalidOperationServerVariables(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Production",
		},
	}

	// Pass invalid type for ContextOperationServerVariables
	ctx := context.WithValue(context.Background(), ContextOperationServerVariables, "invalid")
	_, err := cfg.ServerURLWithContext(ctx, "SomeEndpoint")

	assert.Error(t, err)
}

func TestServerURLWithContext_OperationOverridesGlobal(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Global Server 1",
		},
		{
			URL:         "https://staging.sawvant.com",
			Description: "Global Server 2",
		},
	}
	cfg.OperationServers = map[string]ServerConfigurations{
		"JobsAPIService.GetJob": ServerConfigurations{
			{
				URL:         "http://localhost:9000",
				Description: "Operation-specific server",
			},
		},
	}

	// With operation server configured, it should use that instead of global
	operationIndices := map[string]int{
		"JobsAPIService.GetJob": 0,
	}
	ctx := context.WithValue(context.Background(), ContextOperationServerIndices, operationIndices)
	url, err := cfg.ServerURLWithContext(ctx, "JobsAPIService.GetJob")

	assert.NoError(t, err)
	assert.Equal(t, "http://localhost:9000", url)
}

func TestServerURLWithContext_FallbackToGlobalWhenOperationNotConfigured(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL:         "https://api.sawvant.com",
			Description: "Global Server",
		},
	}
	cfg.OperationServers = map[string]ServerConfigurations{
		"JobsAPIService.GetJob": ServerConfigurations{
			{
				URL:         "http://localhost:9000",
				Description: "Operation-specific",
			},
		},
	}

	// For endpoint not in OperationServers, fall back to global
	url, err := cfg.ServerURLWithContext(nil, "UnknownEndpoint")

	assert.NoError(t, err)
	assert.Equal(t, "https://api.sawvant.com", url)
}

func TestConfiguration_SetDebug(t *testing.T) {
	cfg := NewConfiguration()
	assert.False(t, cfg.Debug)

	cfg.Debug = true
	assert.True(t, cfg.Debug)

	cfg.Debug = false
	assert.False(t, cfg.Debug)
}

func TestConfiguration_SetUserAgent(t *testing.T) {
	cfg := NewConfiguration()
	assert.Equal(t, "sawvant-go-sdk/1.0.0", cfg.UserAgent)

	cfg.UserAgent = "my-app/1.0"
	assert.Equal(t, "my-app/1.0", cfg.UserAgent)
}

func TestConfiguration_SetHostAndScheme(t *testing.T) {
	cfg := NewConfiguration()
	assert.Equal(t, "", cfg.Host)
	assert.Equal(t, "", cfg.Scheme)

	cfg.Host = "api.example.com"
	cfg.Scheme = "http"

	assert.Equal(t, "api.example.com", cfg.Host)
	assert.Equal(t, "http", cfg.Scheme)
}

func TestConfiguration_MultipleDefaultHeaders(t *testing.T) {
	cfg := NewConfiguration()

	cfg.AddDefaultHeader("Authorization", "Bearer token123")
	cfg.AddDefaultHeader("X-API-Key", "key456")
	cfg.AddDefaultHeader("User-Agent", "custom-agent")

	assert.Equal(t, 3, len(cfg.DefaultHeader))
	assert.Equal(t, "Bearer token123", cfg.DefaultHeader["Authorization"])
	assert.Equal(t, "key456", cfg.DefaultHeader["X-API-Key"])
	assert.Equal(t, "custom-agent", cfg.DefaultHeader["User-Agent"])
}

func TestServerConfigurations_URL_ValidIndex(t *testing.T) {
	servers := ServerConfigurations{
		{
			URL:         "https://server1.com",
			Description: "Server 1",
		},
		{
			URL:         "https://server2.com",
			Description: "Server 2",
		},
	}

	url, err := servers.URL(0, nil)
	assert.NoError(t, err)
	assert.Equal(t, "https://server1.com", url)

	url, err = servers.URL(1, nil)
	assert.NoError(t, err)
	assert.Equal(t, "https://server2.com", url)
}

func TestServerConfigurations_URL_InvalidIndex(t *testing.T) {
	servers := ServerConfigurations{
		{
			URL:         "https://server1.com",
			Description: "Server 1",
		},
	}

	_, err := servers.URL(-1, nil)
	assert.Error(t, err)

	_, err = servers.URL(5, nil)
	assert.Error(t, err)
}

func TestServerConfigurations_URL_WithVariableSubstitution(t *testing.T) {
	servers := ServerConfigurations{
		{
			URL:         "https://{protocol}://{host}:{port}",
			Description: "Configurable",
			Variables: map[string]ServerVariable{
				"protocol": {
					Description:  "Protocol",
					DefaultValue: "https",
					EnumValues:   []string{"http", "https"},
				},
				"host": {
					Description:  "Hostname",
					DefaultValue: "api.sawvant.com",
					EnumValues:   []string{"api.sawvant.com", "localhost"},
				},
				"port": {
					Description:  "Port",
					DefaultValue: "443",
					EnumValues:   []string{"443", "8080"},
				},
			},
		},
	}

	// With all defaults
	url, err := servers.URL(0, nil)
	assert.NoError(t, err)
	// Note: URL format in this test is intentionally odd to test substitution
	assert.Contains(t, url, "api.sawvant.com")

	// Override variables
	variables := map[string]string{
		"protocol": "http",
		"host":     "localhost",
		"port":     "8080",
	}
	url, err = servers.URL(0, variables)
	assert.NoError(t, err)
	assert.Contains(t, url, "localhost")
	assert.Contains(t, url, "8080")
}

func TestServerConfigurations_URL_InvalidEnumValue(t *testing.T) {
	servers := ServerConfigurations{
		{
			URL:         "https://{env}.sawvant.com",
			Description: "Environment-specific",
			Variables: map[string]ServerVariable{
				"env": {
					Description:  "Environment",
					DefaultValue: "api",
					EnumValues:   []string{"api", "staging"},
				},
			},
		},
	}

	// Valid enum value
	url, err := servers.URL(0, map[string]string{"env": "staging"})
	assert.NoError(t, err)
	assert.Equal(t, "https://staging.sawvant.com", url)

	// Invalid enum value
	_, err = servers.URL(0, map[string]string{"env": "production"})
	assert.Error(t, err)
}

func TestNewConfiguration_DefaultServerConfiguration(t *testing.T) {
	cfg := NewConfiguration()

	require.NotNil(t, cfg.Servers)
	assert.Equal(t, 1, len(cfg.Servers))
	assert.Equal(t, "https://api.sawvant.com", cfg.Servers[0].URL)
	assert.Equal(t, "Production", cfg.Servers[0].Description)
}

func TestConfiguration_OperationServersInitialized(t *testing.T) {
	cfg := NewConfiguration()

	require.NotNil(t, cfg.OperationServers)
	assert.Equal(t, 0, len(cfg.OperationServers))
}
