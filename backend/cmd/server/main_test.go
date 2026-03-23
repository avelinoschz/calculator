package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnvFloat(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		envValue string
		fallback float64
		want     float64
	}{
		{
			name:     "returns fallback when env var is not set",
			key:      "TEST_PARSE_EMPTY",
			envValue: "",
			fallback: math.Inf(-1),
			want:     math.Inf(-1),
		},
		{
			name:     "returns parsed value when env var is valid",
			key:      "TEST_PARSE_VALID",
			envValue: "42.5",
			fallback: 0,
			want:     42.5,
		},
		{
			name:     "returns fallback when env var is invalid",
			key:      "TEST_PARSE_INVALID",
			envValue: "not-a-number",
			fallback: 100,
			want:     100,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.envValue != "" {
				t.Setenv(tc.key, tc.envValue)
			}
			got := parseEnvFloat(tc.key, tc.fallback)
			if math.IsInf(tc.want, -1) {
				assert.True(t, math.IsInf(got, -1))
			} else {
				assert.InDelta(t, tc.want, got, 1e-9)
			}
		})
	}
}
