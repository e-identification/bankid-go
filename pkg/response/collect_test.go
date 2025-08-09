package response

import (
	"encoding/json"
	"testing"
	"time"
)

func TestBankIdIssueDate_Unmarshal_VariousFormats(t *testing.T) {
	tests := []struct {
		name     string
		jsonIn   string
		wantTime time.Time
	}{
		{
			name:     "RFC3339",
			jsonIn:   `"2025-08-09T12:34:56Z"`,
			wantTime: time.Date(2025, 8, 9, 12, 34, 56, 0, time.UTC),
		},
		{
			name:     "RFC3339Nano",
			jsonIn:   `"2025-08-09T12:34:56.123456789Z"`,
			wantTime: time.Date(2025, 8, 9, 12, 34, 56, 123456789, time.UTC),
		},
		{
			name:     "DateOnlyZ",
			jsonIn:   `"2025-08-09Z"`,
			wantTime: time.Date(2025, 8, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "DateOnlyOffsetUTC",
			jsonIn:   `"2025-08-09+00:00"`,
			wantTime: time.Date(2025, 8, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "DateOnlyOffsetNonUTC",
			jsonIn:   `"2025-08-09+02:00"`,
			wantTime: time.Date(2025, 8, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "DateOnlyNoZone",
			jsonIn:   `"2025-08-09"`,
			wantTime: time.Date(2025, 8, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:     "Null",
			jsonIn:   `null`,
			wantTime: time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d BankIdIssueDate
			if err := json.Unmarshal([]byte(tt.jsonIn), &d); err != nil {
				t.Fatalf("UnmarshalJSON error = %v", err)
			}

			got := d.Time()
			if !got.Equal(tt.wantTime) {
				t.Fatalf("parsed time mismatch: got %s, want %s",
					got.Format(time.RFC3339Nano), tt.wantTime.Format(time.RFC3339Nano))
			}
			// Ensure result is in UTC for non-zero times
			if !got.IsZero() && got.Location() != time.UTC {
				t.Fatalf("time location is not UTC: %v", got.Location())
			}
		})
	}
}

func TestBankIdIssueDate_Marshal_ZeroEmitsNull(t *testing.T) {
	var d BankIdIssueDate

	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("MarshalJSON error = %v", err)
	}

	if string(b) != "null" {
		t.Fatalf("MarshalJSON = %s, want null", string(b))
	}
}

func TestBankIdIssueDate_Marshal_EmitsRFC3339UTC(t *testing.T) {
	// Initialize via Unmarshal (with a non-UTC offset) then Marshal to ensure UTC normalization.
	var d BankIdIssueDate

	in := `"2025-08-09T01:02:03+02:00"`
	if err := json.Unmarshal([]byte(in), &d); err != nil {
		t.Fatalf("UnmarshalJSON error = %v", err)
	}

	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("MarshalJSON error = %v", err)
	}

	// Expect UTC RFC3339 output
	want := `"2025-08-08T23:02:03Z"`
	if string(b) != want {
		t.Fatalf("MarshalJSON = %s, want %s", string(b), want)
	}
}

func TestBankIdIssueDate_RoundTrip_DateOnlyZ(t *testing.T) {
	var d BankIdIssueDate
	if err := json.Unmarshal([]byte(`"2025-08-09Z"`), &d); err != nil {
		t.Fatalf("UnmarshalJSON error = %v", err)
	}

	got := d.Time()

	want := time.Date(2025, 8, 9, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("parsed time mismatch: got %s, want %s", got, want)
	}

	b, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("MarshalJSON error = %v", err)
	}
	// Marshals to RFC3339 in UTC at midnight
	if string(b) != `"2025-08-09T00:00:00Z"` {
		t.Fatalf("MarshalJSON = %s, want %s", string(b), `"2025-08-09T00:00:00Z"`)
	}
}
