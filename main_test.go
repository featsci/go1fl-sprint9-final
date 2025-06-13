package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	cases := []struct {
		desc     string
		vol      int
		expected int
	}{
		{"positive", 2, 2},
		{"with zero", 0, 0},
		{"negative", -1, 0},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			got := generateRandomElements(tt.vol)
			require.Equal(t, len(got), tt.expected)

		})
	}
}

func TestMaximum(t *testing.T) {
	cases := []struct {
		desc     string
		maximum  []int
		expected int
	}{
		{"positive", []int{1, 2, 3}, 3},
		{"with empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"nil", nil, 0},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			got := maximum(tt.maximum)
			require.Equal(t, got, tt.expected)

		})
	}
}
