package utils

import "testing"

func TestParseString(t *testing.T) {
	tests := []struct {
		v        string
		dv       string
		want     string
		testCase string
	}{
		{"hello", "world", "hello", "v is not empty, should return v"},
		{"", "default", "default", "v is empty, should return dv"},
		{"", "", "", "both v and dv are empty, should return empty string"},
	}

	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			got := parseString(tt.v, tt.dv)
			if got != tt.want {
				t.Errorf("parseString(%q, %q) = %q; want %q", tt.v, tt.dv, got, tt.want)
			}
		})
	}
}
