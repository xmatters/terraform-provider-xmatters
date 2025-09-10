package helpers

import (
	"testing"

	"github.com/google/uuid"
)

func TestValidUUID(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    *string
		expected bool
	}{
		{
			name:     "My test case",
			input:    StringPointer(uuid.New().String()),
			expected: true,
		},
		{
			name:     "Valid UUID",
			input:    StringPointer("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11"),
			expected: true,
		},
		{
			name:     "Invalid UUID",
			input:    StringPointer("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11-"),
			expected: false,
		},
		{
			name:     "Empty UUID",
			input:    StringPointer(""),
			expected: false,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidUUID(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
