package main

import "testing"

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		expected string
	}{
		// LeetCode examples

		{
			name:     "example 1",
			n:        3,
			k:        3,
			expected: "213",
		},
		{
			name:     "example 2",
			n:        4,
			k:        9,
			expected: "2314",
		},
		{
			name:     "example 3",
			n:        3,
			k:        1,
			expected: "123",
		},

		// Boundary cases

		{
			name:     "single element",
			n:        1,
			k:        1,
			expected: "1",
		},
		{
			name:     "first permutation",
			n:        4,
			k:        1,
			expected: "1234",
		},
		{
			name:     "last permutation",
			n:        4,
			k:        24,
			expected: "4321",
		},

		// Additional cases

		{
			name:     "second permutation",
			n:        3,
			k:        2,
			expected: "132",
		},
		{
			name:     "middle permutation",
			n:        4,
			k:        12,
			expected: "2431",
		},
		{
			name:     "largest supported n",
			n:        9,
			k:        362880,
			expected: "987654321",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPermutation(tt.n, tt.k)
			if got != tt.expected {
				t.Errorf(
					"GetPermutation(%d, %d) = %q, want %q",
					tt.n, tt.k, got, tt.expected,
				)
			}
		})
	}
}
