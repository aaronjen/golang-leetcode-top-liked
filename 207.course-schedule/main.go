package main

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	preCourseNumTable := make([]int, numCourses)
	nextCourseListTable := make([][]int, numCourses)

	for _, pre := range prerequisites {
		a, b := pre[0], pre[1]
		preCourseNumTable[a]++
		if nextCourseListTable[b] == nil {
			nextCourseListTable[b] = make([]int, 0)
		}
		nextCourseListTable[b] = append(nextCourseListTable[b], a)
	}

	canTakes := []int{}
	for ind, num := range preCourseNumTable {
		if num == 0 {
			canTakes = append(canTakes, ind)
		}
	}

	for len(canTakes) != 0 {
		tmp := []int{}
		for _, course := range canTakes {
			numCourses--

			nextCourses := nextCourseListTable[course]
			if nextCourses != nil {
				for _, nextCourse := range nextCourses {
					preCourseNumTable[nextCourse]--
					if preCourseNumTable[nextCourse] == 0 {
						tmp = append(tmp, nextCourse)
					}
				}
			}
		}
		canTakes = tmp
	}

	return numCourses == 0
}

// @lc code=end
