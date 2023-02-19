package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if index, found := m[target-num]; found && index != i {
			return []int{i, index}
		}
	}
	return []int{}
}
