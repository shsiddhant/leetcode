// Leetcode problem 3609: [Minimum Moves to Reach Target in Grid](https://leetcode.com/problems/minimum-moves-to-reach-target-in-grid)

// You are given four non-negative integers sx, sy, tx, and ty,
// representing two points (sx, sy) and (tx, ty) on an infinitely large 2D grid.

// You start at (sx, sy).

// At any point (x, y), define m = max(x, y). You can either:

// Move to (x + m, y), or
// Move to (x, y + m).

// Return the minimum number of moves required to reach (tx, ty).
// If it is impossible to reach the target, return -1.

// Example 1:
// Input: sx = 1, sy = 2, tx = 5, ty = 4
// Output: 2

// Example 2:
// Input: sx = 0, sy = 1, tx = 2, ty = 3
// Output: 3

// Example 3:
// Input: sx = 1, sy = 1, tx = 2, ty = 2
// Output: -1

// Constraints:

//     0 <= sx <= tx <= 10^9
//     0 <= sy <= ty <= 10^9

package main

import (
	"fmt"
)

// predecessor returns the unique predecessor of (x, y), if it exists.
func predecessor(x, y int) (int, int, error) {
	if x < y {
		if y <= 2*x {
			return x, y - x, nil
		} else if y%2 == 0 {
			return x, y / 2, nil
		} else {
			return 0, 0, fmt.Errorf("no predecessor exists")
		}
	}
	if x > y {
		if x <= 2*y {
			return x - y, y, nil
		} else if x%2 == 0 {
			return x / 2, y, nil
		} else {
			return 0, 0, fmt.Errorf("no predecessor exists")
		}
	}
	return 0, 0, fmt.Errorf("diagonal points has two predecessors")
}

// handleDiagonal handles the case where the target is (t, t).
// A diagonal point (t, t) has exactly two predecessors:
//
//	(0, t) and (t, 0)
func handleDiagonal(sx, sy, t int) int {
	if sx == t && sy == t {
		return 0
	}
	if t == 0 {
		return -1
	}

	if sx == 0 {
		subMoves := axisMoves(sy, t)
		if subMoves != -1 {
			return subMoves + 1
		}
		return -1
	}
	if sy == 0 {
		subMoves := axisMoves(sx, t)
		if subMoves != -1 {
			return subMoves + 1
		}
		return -1
	}
	return -1
}

// axisMoves determines whether an axis point can reach coordinate t.
//
// The returned value is the number of doubling moves required, or -1 if
// s cannot reach t.
func axisMoves(s, t int) int {
	count := 0
	for t > s {
		if t%2 != 0 {
			return -1
		}
		t /= 2
		count++
	}
	if t == s {
		return count
	}
	return -1
}

// NinMoves returns the minimum number of moves required to transform
// (sx, sy) into (tx, ty), or -1 if the target is unreachable.
//
// Transformation rule:
//
// At any point (x, y), define m = max(x, y). You can either:
//
// Move to (x + m, y), or
// Move to (x, y + m).
//
// Assumption: all points have non-negative co-ordinates.
func MinMoves(sx, sy, tx, ty int) int {
	moves := 0
	var err error
	for sx != tx || sy != ty {
		if tx == ty {
			subMoves := handleDiagonal(sx, sy, tx)
			if subMoves == -1 {
				return -1
			} else {
				moves += subMoves
				break
			}
		}
		tx, ty, err = predecessor(tx, ty)
		if err != nil {
			return -1
		}
		moves++
	}
	return moves
}
