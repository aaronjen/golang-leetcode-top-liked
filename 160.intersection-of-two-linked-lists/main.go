package main

/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA := headA
	curB := headB

	for curA != nil || curB != nil {
		if curA == nil {
			curA = headB
		}
		if curB == nil {
			curB = headA
		}
		if curA == curB {
			return curA
		}
		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

// @lc code=end
