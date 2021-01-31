package main

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
func findKthLargest(a []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	k = len(a) - k
	lo, hi := 0, len(a)-1

	for lo < hi {
		j := partition(a, lo, hi)
		if j < k {
			lo = j + 1
		} else if j > k {
			hi = j - 1
		} else {
			break
		}
	}

	return a[k]
}

func partition(nums []int, lo, hi int) int {
	swap := func(a, b int) {
		nums[a], nums[b] = nums[b], nums[a]
	}

	pivot := nums[hi]
	ind := lo
	for i := ind; i < hi; i++ {
		if nums[i] < pivot {
			swap(ind, i)
			ind++
		}
	}
	swap(ind, hi)
	return ind
}

// @lc code=end
