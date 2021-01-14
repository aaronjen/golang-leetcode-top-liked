package main

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	res := head
	for res.Next != nil {
		res, res.Next, prev = res.Next, prev, res
	}
	res.Next = prev

	return res
}

// @lc code=end
