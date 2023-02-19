package main

func findContinuousSequence(target int) [][]int {
	l, r := 1, 1
	sum := r
	var resList [][]int
	for l <= target/2 {
		// move to right
		if sum < target {
			r++
			sum += r
			// move to the left
		} else if sum > target {
			sum -= l
			l++
		} else {
			var res []int
			for i := l; i <= r; i++ {
				res = append(res, i)
			}
			resList = append(resList, res)
			sum -= l
			l++
		}
	}
	return resList

}

/*
func findContinuousSequence(target int) [][]int {
	var res [][]int
	preSum := make([]int, target/2+2)
	for i := 1; i <= target/2+1; i++ {
		preSum[i] = preSum[i-1] + i
	}

	fmt.Println(preSum)
	for i := 2; i < target/2+2; i++ {
		for j := 0; j < i; j++ {
			//fmt.Println(i, j)
			if preSum[i]-preSum[j] == target {
				//fmt.Println("==target")
				var nums []int
				for k := j + 1; k <= i; k++ {
					nums = append(nums, k)
				}
				res = append(res, nums)
			}
		}
		//fmt.Println(res)
	}
	return res
}
*/
