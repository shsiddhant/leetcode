package main

import "testing"

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name     string
		sx, sy   int
		tx, ty   int
		expected int
	}{
		// LeetCode examples

		{
			name: "example 1",
			sx:   1, sy: 2,
			tx: 5, ty: 4,
			expected: 2,
		},
		{
			name: "example 2",
			sx:   0, sy: 1,
			tx: 2, ty: 3,
			expected: 3,
		},
		{
			name: "example 3",
			sx:   1, sy: 1,
			tx: 2, ty: 2,
			expected: -1,
		},

		// Boundary cases

		{
			name: "source equals target",
			sx:   1, sy: 2,
			tx: 1, ty: 2,
			expected: 0,
		},
		{
			name: "origin to origin",
			sx:   0, sy: 0,
			tx: 0, ty: 0,
			expected: 0,
		},
		{
			name: "origin to non-origin",
			sx:   0, sy: 0,
			tx: 1, ty: 1,
			expected: -1,
		},

		// Diagonal cases

		{
			name: "diagonal target unreachable",
			sx:   1, sy: 1,
			tx: 2, ty: 2,
			expected: -1,
		},
		{
			name: "diagonal target from x axis",
			sx:   2, sy: 0,
			tx: 2, ty: 2,
			expected: 1,
		},
		{
			name: "diagonal target from y axis",
			sx:   0, sy: 2,
			tx: 2, ty: 2,
			expected: 1,
		},

		{
			name: "large power of two",
			sx:   1, sy: 1,
			tx: 1, ty: 1 << 29,
			expected: 29,
		},
		{
			name: "greenbutterfly and themonster",
			sx:   192, sy: 1013,
			tx: 1205, ty: 1013,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MinMoves(tt.sx, tt.sy, tt.tx, tt.ty)
			if got != tt.expected {
				t.Errorf(
					"MinMoves(%d, %d, %d, %d) = %d, want %d",
					tt.sx, tt.sy, tt.tx, tt.ty, got, tt.expected,
				)
			}
		})
	}
}
