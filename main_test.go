package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5}, // go test -run '^TestAdd_TableDriven$/positive' ./...
		{"with zero", 0, 5, 5},
		{"negative", -1, -2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			assert.Equalf(t, tt.expected, got, "Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5}, // go test -run '^TestAdd_TableDriven$/positive' ./...
		{"with zero", 0, 5, 5},
		{"negative", -1, -2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			assert.Equalf(t, tt.expected, got, "Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
		})
	}
}

/*
func Add(a, b int) int {
	return a + b
}

func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive", 2, 3, 5}, // go test -run '^TestAdd_TableDriven$/positive' ./...
		{"with zero", 0, 5, 5},
		{"negative", -1, -2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			assert.Equalf(t, tt.expected, got, "Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
		})
	}
}

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if got != want {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}
*/
