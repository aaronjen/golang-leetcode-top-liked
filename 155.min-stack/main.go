package main

/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start

// MinStack struct
type MinStack struct {
	min []int
	all []int
}

// Constructor func
/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		min: []int{},
		all: []int{},
	}
}

// Push func
func (s *MinStack) Push(x int) {
	s.all = append(s.all, x)
	m := s.min
	if len(s.min) == 0 || m[len(m)-1] >= x {
		s.min = append(s.min, x)
	}
}

// Pop func
func (s *MinStack) Pop() {
	all := s.all
	m := s.min
	x := all[len(all)-1]
	s.all = s.all[:len(all)-1]
	if x == m[len(m)-1] {
		s.min = s.min[:len(m)-1]
	}
}

// Top func
func (s *MinStack) Top() int {
	return s.all[len(s.all)-1]
}

// GetMin func
func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
