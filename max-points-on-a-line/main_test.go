package main

import "testing"

func TestReducedSlope(t *testing.T) {
	tests := []struct {
		name string
		p    Vector
		q    Vector
		want Vector
	}{
		{
			name: "same point",
			p:    Vector{1, 2},
			q:    Vector{1, 2},
			want: Vector{0, 0},
		},
		{
			name: "horizontal right",
			p:    Vector{0, 0},
			q:    Vector{5, 0},
			want: Vector{1, 0},
		},
		{
			name: "horizontal left",
			p:    Vector{5, 0},
			q:    Vector{0, 0},
			want: Vector{1, 0},
		},
		{
			name: "horizontal negative coordinates",
			p:    Vector{-10, 5},
			q:    Vector{-2, 5},
			want: Vector{1, 0},
		},
		{
			name: "vertical up",
			p:    Vector{0, 0},
			q:    Vector{0, 5},
			want: Vector{0, 1},
		},
		{
			name: "vertical down",
			p:    Vector{0, 5},
			q:    Vector{0, 0},
			want: Vector{0, 1},
		},
		{
			name: "vertical negative coordinates",
			p:    Vector{-5, -10},
			q:    Vector{-5, -2},
			want: Vector{0, 1},
		},
		{
			name: "positive slope reduced",
			p:    Vector{0, 0},
			q:    Vector{6, 9},
			want: Vector{2, 3},
		},
		{
			name: "negative slope reduced",
			p:    Vector{0, 0},
			q:    Vector{6, -9},
			want: Vector{2, -3},
		},
		{
			name: "reverse positive slope",
			p:    Vector{6, 9},
			q:    Vector{0, 0},
			want: Vector{2, 3},
		},
		{
			name: "reverse negative slope",
			p:    Vector{6, -9},
			q:    Vector{0, 0},
			want: Vector{2, -3},
		},
		{
			name: "already reduced",
			p:    Vector{2, 3},
			q:    Vector{7, 8},
			want: Vector{1, 1},
		},
		{
			name: "gcd greater than two",
			p:    Vector{0, 0},
			q:    Vector{20, 30},
			want: Vector{2, 3},
		},
		{
			name: "gcd with negative dx",
			p:    Vector{10, 10},
			q:    Vector{4, 1},
			want: Vector{2, 3},
		},
		{
			name: "both deltas negative",
			p:    Vector{10, 10},
			q:    Vector{4, 4},
			want: Vector{1, 1},
		},
		{
			name: "negative dx positive dy",
			p:    Vector{5, 1},
			q:    Vector{1, 5},
			want: Vector{1, -1},
		},
		{
			name: "large reduction",
			p:    Vector{-100, -200},
			q:    Vector{200, 400},
			want: Vector{1, 2},
		},
		{
			name: "unit x positive",
			p:    Vector{0, 0},
			q:    Vector{1, 100},
			want: Vector{1, 100},
		},
		{
			name: "unit x negative",
			p:    Vector{0, 0},
			q:    Vector{1, -100},
			want: Vector{1, -100},
		},
		{
			name: "unit y positive",
			p:    Vector{0, 0},
			q:    Vector{100, 1},
			want: Vector{100, 1},
		},
		{
			name: "unit y negative",
			p:    Vector{0, 0},
			q:    Vector{100, -1},
			want: Vector{100, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReducedSlope(tt.p, tt.q)
			if got != tt.want {
				t.Errorf(
					"ReducedSlope(%v, %v) = %v, want %v",
					tt.p, tt.q, got, tt.want,
				)
			}
		})
	}
}

func TestMaxCollinear(t *testing.T) {
	tests := []struct {
		name   string
		points []Vector
		want   int
	}{
		// Leetcode
		{
			name: "example 1",
			points: []Vector{
				{1, 1},
				{2, 2},
				{3, 3},
			},
			want: 3,
		},
		{
			name: "example 2",
			points: []Vector{
				{1, 1},
				{3, 2},
				{5, 3},
				{4, 1},
				{2, 3},
				{1, 4},
			},
			want: 4,
		},
		// Additonal tests
		{
			name:   "nil",
			points: nil,
			want:   0,
		},
		{
			name:   "empty",
			points: []Vector{},
			want:   0,
		},
		{
			name:   "one point",
			points: []Vector{{0, 0}},
			want:   1,
		},
		{
			name: "two distinct points",
			points: []Vector{
				{0, 0},
				{1, 1},
			},
			want: 2,
		},
		{
			name: "two identical points",
			points: []Vector{
				{0, 0},
				{0, 0},
			},
			want: 2,
		},
		{
			name: "all points identical",
			points: []Vector{
				{5, 5},
				{5, 5},
				{5, 5},
				{5, 5},
			},
			want: 4,
		},
		{
			name: "all points horizontal",
			points: []Vector{
				{0, 5},
				{1, 5},
				{2, 5},
				{10, 5},
			},
			want: 4,
		},
		{
			name: "all points vertical",
			points: []Vector{
				{5, 0},
				{5, 1},
				{5, 2},
				{5, 10},
			},
			want: 4,
		},
		{
			name: "all points positive diagonal",
			points: []Vector{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
				{10, 10},
			},
			want: 5,
		},
		{
			name: "all points negative diagonal",
			points: []Vector{
				{0, 5},
				{1, 4},
				{2, 3},
				{3, 2},
				{5, 0},
			},
			want: 5,
		},
		{
			name: "reduced slope required",
			points: []Vector{
				{0, 0},
				{2, 4},
				{4, 8},
				{6, 12},
				{1, 0},
			},
			want: 4,
		},
		{
			name: "no three collinear",
			points: []Vector{
				{0, 0},
				{1, 2},
				{2, 1},
				{3, 5},
			},
			want: 2,
		},
		{
			name: "maximum is horizontal",
			points: []Vector{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
				{0, 1},
				{1, 2},
			},
			want: 4,
		},
		{
			name: "maximum is vertical",
			points: []Vector{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},
				{1, 1},
				{2, 2},
			},
			want: 4,
		},
		{
			name: "maximum is positive diagonal",
			points: []Vector{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
				{4, 4},
				{0, 1},
				{0, 2},
			},
			want: 5,
		},
		{
			name: "maximum is negative diagonal",
			points: []Vector{
				{0, 5},
				{1, 4},
				{2, 3},
				{3, 2},
				{4, 1},
				{0, 0},
				{1, 1},
			},
			want: 5,
		},
		{
			name: "negative coordinates",
			points: []Vector{
				{-3, -3},
				{-2, -2},
				{-1, -1},
				{0, 0},
				{1, 1},
			},
			want: 5,
		},
		{
			name: "mixed positive and negative coordinates",
			points: []Vector{
				{-2, -4},
				{-1, -2},
				{0, 0},
				{1, 2},
				{2, 4},
				{10, 10},
			},
			want: 5,
		},
		{
			name: "two equally large lines",
			points: []Vector{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},

				{0, 10},
				{1, 9},
				{2, 8},
				{3, 7},
			},
			want: 4,
		},
		{
			name: "vertical and horizontal candidates",
			points: []Vector{
				{0, 0},
				{0, 1},
				{0, 2},
				{0, 3},

				{1, 0},
				{2, 0},
				{3, 0},
			},
			want: 4,
		},
		{
			name: "one point lies on several candidate lines",
			points: []Vector{
				{0, 0},
				{1, 1},
				{2, 2},
				{3, 3},
				{1, 0},
				{2, 0},
				{3, 0},
			},
			want: 4,
		},
		{
			name: "large coordinate differences",
			points: []Vector{
				{-100, -200},
				{0, 0},
				{100, 200},
				{200, 400},
				{300, 600},
			},
			want: 5,
		},
		{
			name: "large coordinates with reduced slope",
			points: []Vector{
				{-100000, -200000},
				{0, 0},
				{100000, 200000},
				{200000, 400000},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxCollinear(tt.points)
			if got != tt.want {
				t.Errorf(
					"MaxCollinear(%v) = %d, want %d",
					tt.points, got, tt.want,
				)
			}
		})
	}
}
