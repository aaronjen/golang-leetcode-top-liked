package main

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := split(head)
	if left != nil {
		left = sortList(left)
	}
	if right != nil {
		right = sortList(right)
	}
	return merge(left, right)
}

func split(node *ListNode) (*ListNode, *ListNode) {
	fast := node
	slow := node
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slow.Next, slow = nil, slow.Next
	return node, slow
}

func merge(a, b *ListNode) *ListNode {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val > b.Val {
		b.Next = merge(a, b.Next)
		return b
	}
	a.Next = merge(a.Next, b)
	return a
}

// @lc code=end
