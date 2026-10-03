# Minimum Moves to Reach Target in Grid

A Go implementation of the leetcode problem [3609](https://leetcode.com/problems/minimum-moves-to-reach-target-in-grid), along with the mathematics behind the algorithm.

## Problem

You are given four non-negative integers `sx`, `sy`, `tx`, and `ty`, representing two points `(sx, sy)` and `(tx, ty)` on an infinitely large 2D grid.

You start at `(sx, sy)`.

At any point `(x, y)`, define `m = max(x, y)`. You can either:

- Move to `(x + m, y)`, or
- Move to `(x, y + m)`.

Return the minimum number of moves required to reach `(tx, ty)`. If it is impossible to reach the target, return `-1`.

## Approach

The main insight behind the algorithm is this: instead of searching for the target in a sea of forward paths, we can work backwards from the target.

For most points, there is at most one possible predecessor. This makes the reverse search essentially deterministic. The only points where a choice can occur are points on the diagonal (x, x), and the source determines which branch can lead back to it.

The algorithm therefore:

- Starts at the target.
- Finds its possible predecessor.
- Repeats until it reaches the source or determines that no path exists.
- Counts the reverse steps to obtain the minimum number of forward moves.

For more details on the math, see [MATH.md]

## Examples

1. source = `(1, 2)` and target = `(5, 8)`

```
(5, 8)
   ↓
(5, 3)
   ↓
(2, 3)
   ↓
(1, 2)
```

So `(5, 8)` is reachable from `(1, 2)` in 3 moves.

2. source = `(1,1)`, target = `(2, 2)`

```
(2, 2)
   ↓
(2, 0) or (0, 2)
```

The target lies on the diagonal, so there are two possible predecessors. However, both predecessors lie on an axis. Since source doesn't lie on either axis, neither branch can lead back to it.

Therefore, the target is not reachable, and the algorithm returns `-1`.

## Complexity

Let `M = max(tx, ty)`.

The reverse algorithm alternates between subtracting the smaller
coordinate and halving the larger coordinate. Every two reverse steps
reduce the sum of the coordinates by at least a factor of two.

Therefore:

- **Time:** `O(log M)`
- **Space:** `O(1)`
