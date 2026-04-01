package sawvant

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptimizeRequest_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		request  *OptimizeRequest
		validate func(t *testing.T, data []byte)
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: NewOptimizeRequest(
				[]Part{
					{Id: "part-1", Length: 100, Width: 50, Quantity: 1, Grain: LENGTH},
				},
				[]Sheet{
					{Id: "sheet-1", Length: 1000, Width: 500, Quantity: 10, Grain: LENGTH},
				},
				Machine{BladeThickness: 3.5, MaxLevels: 2, CutDirection: DEFAULT},
			),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.NotNil(t, m["parts"])
				assert.NotNil(t, m["sheets"])
				assert.NotNil(t, m["machine"])
				// strategy has a default value
				assert.NotNil(t, m["strategy"])
				assert.Equal(t, "fast", m["strategy"])
				// cost_tariffs should not be present when nil
				assert.Nil(t, m["cost_tariffs"])
			},
			wantErr: false,
		},
		{
			name: "with optional cost_tariffs",
			request: func() *OptimizeRequest {
				req := NewOptimizeRequest(
					[]Part{{Id: "part-1", Length: 100, Width: 50, Quantity: 1, Grain: WIDTH}},
					[]Sheet{{Id: "sheet-1", Length: 1000, Width: 500, Quantity: 10, Grain: WIDTH}},
					Machine{BladeThickness: 3.5, MaxLevels: 2, CutDirection: RIP},
				)
				req.CostTariffs = &CostTariffs{}
				return req
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.NotNil(t, m["cost_tariffs"])
			},
			wantErr: false,
		},
		{
			name: "with custom strategy",
			request: func() *OptimizeRequest {
				req := NewOptimizeRequest(
					[]Part{{Id: "part-1", Length: 100, Width: 50, Quantity: 1, Grain: NONE}},
					[]Sheet{{Id: "sheet-1", Length: 1000, Width: 500, Quantity: 10, Grain: NONE}},
					Machine{BladeThickness: 3.5, MaxLevels: 2, CutDirection: CROSS},
				)
				strategy := "thorough"
				req.Strategy = &strategy
				return req
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "thorough", m["strategy"])
			},
			wantErr: false,
		},
		{
			name: "empty parts and sheets",
			request: NewOptimizeRequest(
				[]Part{},
				[]Sheet{},
				Machine{BladeThickness: 3.5, MaxLevels: 2, CutDirection: DEFAULT},
			),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				parts := m["parts"].([]interface{})
				sheets := m["sheets"].([]interface{})
				assert.Equal(t, 0, len(parts))
				assert.Equal(t, 0, len(sheets))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.request)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tt.validate(t, data)
			}
		})
	}
}

func TestOptimizeRequest_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		validate func(t *testing.T, req *OptimizeRequest)
	}{
		{
			name: "valid request with required fields",
			json: `{
				"parts": [{"id":"p1","length":100,"width":50,"quantity":1,"grain":"length"}],
				"sheets": [{"id":"s1","length":1000,"width":500,"quantity":10,"grain":"length"}],
				"machine": {"blade_thickness":3.5,"max_levels":2,"cut_direction":"default"}
			}`,
			wantErr: false,
			validate: func(t *testing.T, req *OptimizeRequest) {
				assert.Equal(t, 1, len(req.Parts))
				assert.Equal(t, 1, len(req.Sheets))
				assert.Equal(t, 3.5, req.Machine.BladeThickness)
			},
		},
		{
			name: "missing required field: parts",
			json: `{
				"sheets": [{"id":"s1","length":1000,"width":500,"quantity":10,"grain":"length"}],
				"machine": {"blade_thickness":3.5,"max_levels":2,"cut_direction":"default"}
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: sheets",
			json: `{
				"parts": [{"id":"p1","length":100,"width":50,"quantity":1,"grain":"length"}],
				"machine": {"blade_thickness":3.5,"max_levels":2,"cut_direction":"default"}
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: machine",
			json: `{
				"parts": [{"id":"p1","length":100,"width":50,"quantity":1,"grain":"length"}],
				"sheets": [{"id":"s1","length":1000,"width":500,"quantity":10,"grain":"length"}]
			}`,
			wantErr: true,
		},
		{
			name: "unknown field rejected",
			json: `{
				"parts": [{"id":"p1","length":100,"width":50,"quantity":1,"grain":"length"}],
				"sheets": [{"id":"s1","length":1000,"width":500,"quantity":10,"grain":"length"}],
				"machine": {"blade_thickness":3.5,"max_levels":2,"cut_direction":"default"},
				"unknown_field": "should cause error"
			}`,
			wantErr: true,
		},
		{
			name: "with optional strategy",
			json: `{
				"parts": [{"id":"p1","length":100,"width":50,"quantity":1,"grain":"length"}],
				"sheets": [{"id":"s1","length":1000,"width":500,"quantity":10,"grain":"length"}],
				"machine": {"blade_thickness":3.5,"max_levels":2,"cut_direction":"default"},
				"strategy": "thorough"
			}`,
			wantErr: false,
			validate: func(t *testing.T, req *OptimizeRequest) {
				assert.NotNil(t, req.Strategy)
				assert.Equal(t, "thorough", *req.Strategy)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req OptimizeRequest
			err := json.Unmarshal([]byte(tt.json), &req)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, &req)
				}
			}
		})
	}
}

func TestJobResponse_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		response *JobResponse
		validate func(t *testing.T, data []byte)
	}{
		{
			name: "minimal job response",
			response: NewJobResponse(
				"job-1",
				"pending",
				0,
				time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "job-1", m["job_id"])
				assert.Equal(t, "pending", m["status"])
				assert.Equal(t, float64(0), m["progress"])
				assert.NotNil(t, m["created_at"])
				// Optional fields should not be present
				assert.Nil(t, m["started_at"])
				assert.Nil(t, m["completed_at"])
				assert.Nil(t, m["result"])
				assert.Nil(t, m["warnings"])
				assert.Nil(t, m["error"])
			},
		},
		{
			name: "job response with all fields",
			response: func() *JobResponse {
				resp := NewJobResponse(
					"job-2",
					"completed",
					100,
					time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				)
				started := time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC)
				completed := time.Date(2024, 1, 1, 2, 0, 0, 0, time.UTC)
				resp.StartedAt = &started
				resp.CompletedAt = &completed
				resp.Warnings = []string{"warning-1"}
				resp.Result = &OptimizeResult{}
				return resp
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.NotNil(t, m["started_at"])
				assert.NotNil(t, m["completed_at"])
				assert.NotNil(t, m["warnings"])
				assert.NotNil(t, m["result"])
			},
		},
		{
			name: "job response with error",
			response: func() *JobResponse {
				resp := NewJobResponse(
					"job-3",
					"failed",
					0,
					time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				)
				errorMsg := "Something went wrong"
				resp.Error = &errorMsg
				return resp
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "Something went wrong", m["error"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.response)
			require.NoError(t, err)
			tt.validate(t, data)
		})
	}
}

func TestJobResponse_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		validate func(t *testing.T, resp *JobResponse)
	}{
		{
			name: "valid minimal response",
			json: `{
				"job_id": "job-1",
				"status": "running",
				"progress": 50,
				"created_at": "2024-01-01T00:00:00Z"
			}`,
			wantErr: false,
			validate: func(t *testing.T, resp *JobResponse) {
				assert.Equal(t, "job-1", resp.JobId)
				assert.Equal(t, "running", resp.Status)
				assert.Equal(t, int32(50), resp.Progress)
				assert.Nil(t, resp.StartedAt)
				assert.Nil(t, resp.CompletedAt)
			},
		},
		{
			name: "missing required field: job_id",
			json: `{
				"status": "running",
				"progress": 50,
				"created_at": "2024-01-01T00:00:00Z"
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: status",
			json: `{
				"job_id": "job-1",
				"progress": 50,
				"created_at": "2024-01-01T00:00:00Z"
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: progress",
			json: `{
				"job_id": "job-1",
				"status": "running",
				"created_at": "2024-01-01T00:00:00Z"
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: created_at",
			json: `{
				"job_id": "job-1",
				"status": "running",
				"progress": 50
			}`,
			wantErr: true,
		},
		{
			name: "with optional timestamps",
			json: `{
				"job_id": "job-1",
				"status": "completed",
				"progress": 100,
				"created_at": "2024-01-01T00:00:00Z",
				"started_at": "2024-01-01T01:00:00Z",
				"completed_at": "2024-01-01T02:00:00Z"
			}`,
			wantErr: false,
			validate: func(t *testing.T, resp *JobResponse) {
				assert.NotNil(t, resp.StartedAt)
				assert.NotNil(t, resp.CompletedAt)
			},
		},
		{
			name: "unknown field rejected",
			json: `{
				"job_id": "job-1",
				"status": "running",
				"progress": 50,
				"created_at": "2024-01-01T00:00:00Z",
				"unknown_field": "value"
			}`,
			wantErr: true,
		},
		{
			name: "all statuses",
			json: `{
				"job_id": "job-1",
				"status": "completed",
				"progress": 100,
				"created_at": "2024-01-01T00:00:00Z"
			}`,
			wantErr: false,
			validate: func(t *testing.T, resp *JobResponse) {
				assert.Equal(t, "completed", resp.Status)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp JobResponse
			err := json.Unmarshal([]byte(tt.json), &resp)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, &resp)
				}
			}
		})
	}
}

func TestPart_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		part     *Part
		validate func(t *testing.T, data []byte)
	}{
		{
			name: "basic part with all required fields",
			part: NewPart("part-1", 100.5, 50.25, 5, LENGTH),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "part-1", m["id"])
				assert.Equal(t, 100.5, m["length"])
				assert.Equal(t, 50.25, m["width"])
				assert.Equal(t, float64(5), m["quantity"])
				assert.Equal(t, "length", m["grain"])
				assert.Nil(t, m["edge_banding"])
			},
		},
		{
			name: "part with edge banding",
			part: func() *Part {
				part := NewPart("part-1", 100.5, 50.25, 5, WIDTH)
				part.EdgeBanding = &EdgeCorrection{}
				return part
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.NotNil(t, m["edge_banding"])
			},
		},
		{
			name: "part with grain directions",
			part: NewPart("part-2", 200, 100, 10, NONE),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "none", m["grain"])
			},
		},
		{
			name: "part with free_same grain",
			part: NewPart("part-3", 150, 75, 3, FREE_SAME),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "free_same", m["grain"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.part)
			require.NoError(t, err)
			tt.validate(t, data)
		})
	}
}

func TestPart_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		validate func(t *testing.T, part *Part)
	}{
		{
			name: "valid part with required fields",
			json: `{
				"id": "part-1",
				"length": 100.5,
				"width": 50.25,
				"quantity": 5,
				"grain": "length"
			}`,
			wantErr: false,
			validate: func(t *testing.T, part *Part) {
				assert.Equal(t, "part-1", part.Id)
				assert.Equal(t, 100.5, part.Length)
				assert.Equal(t, 50.25, part.Width)
				assert.Equal(t, int32(5), part.Quantity)
				assert.Equal(t, LENGTH, part.Grain)
			},
		},
		{
			name: "missing required field: id",
			json: `{
				"length": 100.5,
				"width": 50.25,
				"quantity": 5,
				"grain": "length"
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: grain",
			json: `{
				"id": "part-1",
				"length": 100.5,
				"width": 50.25,
				"quantity": 5
			}`,
			wantErr: true,
		},
		{
			name: "all grain directions",
			json: `{
				"id": "part-1",
				"length": 100,
				"width": 50,
				"quantity": 1,
				"grain": "free_same"
			}`,
			wantErr: false,
			validate: func(t *testing.T, part *Part) {
				assert.Equal(t, FREE_SAME, part.Grain)
			},
		},
		{
			name: "unknown field rejected",
			json: `{
				"id": "part-1",
				"length": 100,
				"width": 50,
				"quantity": 1,
				"grain": "length",
				"unknown_field": "value"
			}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var part Part
			err := json.Unmarshal([]byte(tt.json), &part)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, &part)
				}
			}
		})
	}
}

func TestSheet_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		sheet    *Sheet
		validate func(t *testing.T, data []byte)
	}{
		{
			name: "basic sheet with required fields",
			sheet: NewSheet("sheet-1", 1000, 500, 10, LENGTH),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "sheet-1", m["id"])
				assert.Equal(t, float64(1000), m["length"])
				assert.Equal(t, float64(500), m["width"])
				assert.Equal(t, float64(10), m["quantity"])
				assert.Equal(t, "length", m["grain"])
				// Default is_offcut should be false but still present
				assert.NotNil(t, m["is_offcut"])
				assert.False(t, m["is_offcut"].(bool))
			},
		},
		{
			name: "sheet with is_offcut=true",
			sheet: func() *Sheet {
				sheet := NewSheet("sheet-1", 1000, 500, 10, WIDTH)
				isOffcut := true
				sheet.IsOffcut = &isOffcut
				return sheet
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.True(t, m["is_offcut"].(bool))
			},
		},
		{
			name: "sheet with trim margins",
			sheet: func() *Sheet {
				sheet := NewSheet("sheet-1", 1000, 500, 10, NONE)
				sheet.TrimMargins = &Margins{}
				return sheet
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.NotNil(t, m["trim_margins"])
			},
		},
		{
			name: "sheet with article number",
			sheet: func() *Sheet {
				sheet := NewSheet("sheet-1", 1000, 500, 10, LENGTH)
				articleNum := "SKU-12345"
				sheet.ArticleNumber = &articleNum
				return sheet
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "SKU-12345", m["article_number"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.sheet)
			require.NoError(t, err)
			tt.validate(t, data)
		})
	}
}

func TestSheet_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		validate func(t *testing.T, sheet *Sheet)
	}{
		{
			name: "valid sheet with required fields",
			json: `{
				"id": "sheet-1",
				"length": 1000,
				"width": 500,
				"quantity": 10,
				"grain": "length"
			}`,
			wantErr: false,
			validate: func(t *testing.T, sheet *Sheet) {
				assert.Equal(t, "sheet-1", sheet.Id)
				assert.Equal(t, float64(1000), sheet.Length)
				assert.Equal(t, float64(500), sheet.Width)
				assert.Equal(t, int32(10), sheet.Quantity)
				assert.Equal(t, LENGTH, sheet.Grain)
			},
		},
		{
			name: "missing required field: id",
			json: `{
				"length": 1000,
				"width": 500,
				"quantity": 10,
				"grain": "length"
			}`,
			wantErr: true,
		},
		{
			name: "with optional is_offcut",
			json: `{
				"id": "sheet-1",
				"length": 1000,
				"width": 500,
				"quantity": 10,
				"grain": "length",
				"is_offcut": true
			}`,
			wantErr: false,
			validate: func(t *testing.T, sheet *Sheet) {
				assert.NotNil(t, sheet.IsOffcut)
				assert.True(t, *sheet.IsOffcut)
			},
		},
		{
			name: "with trim_margins",
			json: `{
				"id": "sheet-1",
				"length": 1000,
				"width": 500,
				"quantity": 10,
				"grain": "width",
				"trim_margins": {"left": 10}
			}`,
			wantErr: false,
			validate: func(t *testing.T, sheet *Sheet) {
				assert.NotNil(t, sheet.TrimMargins)
			},
		},
		{
			name: "with article_number",
			json: `{
				"id": "sheet-1",
				"length": 1000,
				"width": 500,
				"quantity": 10,
				"grain": "none",
				"article_number": "SKU-12345"
			}`,
			wantErr: false,
			validate: func(t *testing.T, sheet *Sheet) {
				assert.NotNil(t, sheet.ArticleNumber)
				assert.Equal(t, "SKU-12345", *sheet.ArticleNumber)
			},
		},
		{
			name: "unknown field rejected",
			json: `{
				"id": "sheet-1",
				"length": 1000,
				"width": 500,
				"quantity": 10,
				"grain": "length",
				"unknown_field": "value"
			}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sheet Sheet
			err := json.Unmarshal([]byte(tt.json), &sheet)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, &sheet)
				}
			}
		})
	}
}

func TestMachine_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		machine  *Machine
		validate func(t *testing.T, data []byte)
	}{
		{
			name: "basic machine with required fields",
			machine: NewMachine(3.5, 2, DEFAULT),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, 3.5, m["blade_thickness"])
				assert.Equal(t, float64(2), m["max_levels"])
				assert.Equal(t, "default", m["cut_direction"])
				// max_stack_height should not be present when nil
				assert.Nil(t, m["max_stack_height"])
			},
		},
		{
			name: "machine with max_stack_height",
			machine: func() *Machine {
				machine := NewMachine(3.5, 2, RIP)
				stackHeight := 100.0
				machine.MaxStackHeight = &stackHeight
				return machine
			}(),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, 100.0, m["max_stack_height"])
			},
		},
		{
			name: "machine with all cut directions",
			machine: NewMachine(4.0, 3, CROSS),
			validate: func(t *testing.T, data []byte) {
				var m map[string]interface{}
				err := json.Unmarshal(data, &m)
				require.NoError(t, err)

				assert.Equal(t, "cross", m["cut_direction"])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.machine)
			require.NoError(t, err)
			tt.validate(t, data)
		})
	}
}

func TestMachine_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		json    string
		wantErr bool
		validate func(t *testing.T, machine *Machine)
	}{
		{
			name: "valid machine with required fields",
			json: `{
				"blade_thickness": 3.5,
				"max_levels": 2,
				"cut_direction": "default"
			}`,
			wantErr: false,
			validate: func(t *testing.T, machine *Machine) {
				assert.Equal(t, 3.5, machine.BladeThickness)
				assert.Equal(t, int32(2), machine.MaxLevels)
				assert.Equal(t, DEFAULT, machine.CutDirection)
			},
		},
		{
			name: "missing required field: blade_thickness",
			json: `{
				"max_levels": 2,
				"cut_direction": "default"
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: max_levels",
			json: `{
				"blade_thickness": 3.5,
				"cut_direction": "default"
			}`,
			wantErr: true,
		},
		{
			name: "missing required field: cut_direction",
			json: `{
				"blade_thickness": 3.5,
				"max_levels": 2
			}`,
			wantErr: true,
		},
		{
			name: "with optional max_stack_height",
			json: `{
				"blade_thickness": 3.5,
				"max_levels": 2,
				"cut_direction": "rip",
				"max_stack_height": 150.5
			}`,
			wantErr: false,
			validate: func(t *testing.T, machine *Machine) {
				assert.NotNil(t, machine.MaxStackHeight)
				assert.Equal(t, 150.5, *machine.MaxStackHeight)
			},
		},
		{
			name: "all cut directions",
			json: `{
				"blade_thickness": 4.0,
				"max_levels": 3,
				"cut_direction": "cross"
			}`,
			wantErr: false,
			validate: func(t *testing.T, machine *Machine) {
				assert.Equal(t, CROSS, machine.CutDirection)
			},
		},
		{
			name: "unknown field rejected",
			json: `{
				"blade_thickness": 3.5,
				"max_levels": 2,
				"cut_direction": "default",
				"unknown_field": "value"
			}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var machine Machine
			err := json.Unmarshal([]byte(tt.json), &machine)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, &machine)
				}
			}
		})
	}
}

func TestGrainDirection_Enum(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
		validate func(t *testing.T, grain GrainDirection)
	}{
		{
			name:    "none",
			value:   "none",
			wantErr: false,
			validate: func(t *testing.T, grain GrainDirection) {
				assert.Equal(t, NONE, grain)
			},
		},
		{
			name:    "length",
			value:   "length",
			wantErr: false,
			validate: func(t *testing.T, grain GrainDirection) {
				assert.Equal(t, LENGTH, grain)
			},
		},
		{
			name:    "width",
			value:   "width",
			wantErr: false,
			validate: func(t *testing.T, grain GrainDirection) {
				assert.Equal(t, WIDTH, grain)
			},
		},
		{
			name:    "free_same",
			value:   "free_same",
			wantErr: false,
			validate: func(t *testing.T, grain GrainDirection) {
				assert.Equal(t, FREE_SAME, grain)
			},
		},
		{
			name:    "invalid value",
			value:   "invalid_direction",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var grain GrainDirection
			err := json.Unmarshal([]byte(`"`+tt.value+`"`), &grain)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, grain)
				}
			}
		})
	}
}

func TestCutDirection_Enum(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		wantErr bool
		validate func(t *testing.T, dir CutDirection)
	}{
		{
			name:    "default",
			value:   "default",
			wantErr: false,
			validate: func(t *testing.T, dir CutDirection) {
				assert.Equal(t, DEFAULT, dir)
			},
		},
		{
			name:    "rip",
			value:   "rip",
			wantErr: false,
			validate: func(t *testing.T, dir CutDirection) {
				assert.Equal(t, RIP, dir)
			},
		},
		{
			name:    "cross",
			value:   "cross",
			wantErr: false,
			validate: func(t *testing.T, dir CutDirection) {
				assert.Equal(t, CROSS, dir)
			},
		},
		{
			name:    "invalid direction",
			value:   "diagonal",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dir CutDirection
			err := json.Unmarshal([]byte(`"`+tt.value+`"`), &dir)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if tt.validate != nil {
					tt.validate(t, dir)
				}
			}
		})
	}
}
