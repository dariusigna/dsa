package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	mp := make(map[string][]int)
	threeSumUtil(nums, 0, []int{}, mp)
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}

func zeroSum(arr []int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum == 0
}

func createStringID(arr []int) string {
	buffer := strings.Builder{}
	for _, v := range arr {
		buffer.Write([]byte((strconv.Itoa(v))))
	}

	return buffer.String()
}

func threeSumUtil(nums []int, k int, sol []int, result map[string][]int) {
	if len(sol) == 3 && zeroSum(sol) {
		cpy := make([]int, len(sol))
		copy(cpy, sol)
		result[createStringID(sol)] = cpy
		return
	}

	for i := k; i < len(nums); i++ {
		sol = append(sol, nums[i])
		threeSumUtil(nums, i+1, sol, result)
		sol = sol[:len(sol)-1]
	}
}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
