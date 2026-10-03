// Leetcode problem 149: [Max Points on a Line](https://leetcode.com/problems/max-points-on-a-line)

// Given an array of points where points[i] = [x_i, y_i] represents a point on the X-Y plane,
// return the maximum number of points that lie on the same straight line.

// Example 1:
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 3

// Example 2:
// Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4

// Constraints:

//     1 <= points.length <= 300
//     points[i].length == 2
//     -10^4 <= x_i, y_i <= 10^4
//     All the points are unique.

package main

type Vector struct {
	X, Y int
}

func (p Vector) Sub(q Vector) Vector {
	return Vector{p.X - q.X, p.Y - q.Y}
}

// ReducedSlope computes the reduced slope of the line joining two points p & q.
//
// Reduced slope of p and q is defined as the unique vector v satisfying:
//
// 1. v is parallel to p-q
//
// 2. v.X and v.Y are coprime unless one of them is zero
//
// 3. v.X >= 0
//
// 4. If v.X = 0, then v.Y >= 0
func ReducedSlope(p, q Vector) Vector {
	dx := q.X - p.X
	dy := q.Y - p.Y

	if dx == 0 && dy == 0 {
		return Vector{}
	}

	g := abs(gcd(dx, dy))
	dx /= g
	dy /= g

	if dx < 0 || (dx == 0 && dy < 0) {
		dx = -dx
		dy = -dy
	}

	return Vector{dx, dy}
}

// MaxCollinear returns the maximum number of points that lie on the
// same straight line for a given array of points.
func MaxCollinear(points []Vector) int {
	n := len(points)
	if n == 0 {
		return 0
	}
	max := 0
	for i := range n {
		// n-i-1 points to check (i < j < n). If that's <= max, break the loop
		if n-i-1 <= max {
			break
		}
		slopes := make(map[Vector]int)
		for j := i + 1; j < n; j++ {
			slope := ReducedSlope(points[i], points[j])
			slopes[slope]++
			if slopes[slope] > max {
				max = slopes[slope]
			}
		}
	}
	return max + 1
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
