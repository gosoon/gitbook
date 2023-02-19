package main

/*
func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return []int{}
	}
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{}
}
*/

func twoSum(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		if index, found := m[target-numbers[i]]; found {
			return []int{index + 1, i + 1}
		}
		m[numbers[i]] = i

	}
	return []int{}
}
