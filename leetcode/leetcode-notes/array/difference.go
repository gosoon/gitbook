package main

func DifferenceResult(difference []int) []int {
	res := make([]int, len(difference))
	res[0] = difference[0]
	for i := 1; i < len(difference); i++ {
		res[i] = difference[i] + res[i-1]
	}
	return res
}

func DifferenceIncrement(difference []int, i int, j int, val int) []int {
	difference[i] += val
	if len(difference) > j+1 {
		difference[j+1] -= val
	}
	return difference
}
