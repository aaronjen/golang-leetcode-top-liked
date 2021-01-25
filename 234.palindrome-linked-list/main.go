package main

/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	dummy := &ListNode{
		Next: head,
	}
	fast := dummy
	slow := dummy

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := reverseLis(slow.Next)
	slow.Next = nil

	for head != nil && mid != nil {
		if head.Val != mid.Val {
			return false
		}

		head = head.Next
		mid = mid.Next
	}

	return true
}

func reverseLis(head *ListNode) *ListNode {
	res := &ListNode{
		Next: head,
	}

	var prev *ListNode
	for res.Next.Next != nil {
		res.Next, res.Next.Next, prev = res.Next.Next, prev, res.Next
	}
	res.Next.Next = prev

	return res.Next
}

// @lc code=end
