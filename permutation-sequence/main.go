// Leetcode problem 60: [Permutation Sequence](https://leetcode.com/problems/permutation-sequence)

// The set [1, 2, 3, ..., n] contains a total of n! unique permutations.

// By listing and labeling all of the permutations in order, we get the following sequence for n = 3:

//     "123"
//     "132"
//     "213"
//     "231"
//     "312"
//     "321"

// Given n and k, return the kth permutation sequence.

// Example 1:

// Input: n = 3, k = 3
// Output: "213"

// Example 2:

// Input: n = 4, k = 9
// Output: "2314"

// Example 3:

// Input: n = 3, k = 1
// Output: "123"

// Constraints:
//
//	1 <= n <= 9
//	1 <= k <= n!

package main

// GetPermutation returns the k-th permutation in lexicographic order of
// the integers 1 through n.
//
// The permutations are partitioned into blocks according to their first
// element. Each block contains (n-1)! permutations. After selecting the
// first element, the same idea is applied recursively to the remaining
// elements.
//
// Indexing for k starts at 1.
func GetPermutation(n, k int) string {
	results := make([]byte, n)
	used := make(map[int]struct{}, n)
	blockSize := factorial(n - 1)

	for i := range n {
		position := (k-1)/blockSize + 1

		value := skipIndices(position, n, used)
		results[i] = byte('0' + value)
		used[value] = struct{}{}

		k = (k-1)%blockSize + 1

		if i < n-1 {
			blockSize /= n - i - 1
		}
	}

	return string(results)
}

// factorial returns n!
func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// skipIndices returns the position-th unused value between 1 and 'max'.
//
// position is 1-based. Values already present in 'used' are skipped when
// counting positions. For example, if used contains {2}:
//
//	position 1 -> 1
//	position 2 -> 3
//	position 3 -> 4
func skipIndices(position, max int, used map[int]struct{}) int {
	for i := 1; i <= max; i++ {
		if position >= i {
			if _, ok := used[i]; ok {
				position++
			}
		}
	}
	return position
}
