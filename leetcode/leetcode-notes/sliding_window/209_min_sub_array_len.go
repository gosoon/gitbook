package main

func minSubArrayLen(target int, nums []int) int {
	var l, r, sum int
	minSubArrayLen := len(nums) + 1
	for r < len(nums) {
		// scale up windows
		t := nums[r]
		r++
		sum += t

		//fmt.Printf("l:%v,r:%v,result:%v,sum:%v \n", l, r, minSubArrayLen, sum)
		for sum >= target {
			if r-l < minSubArrayLen {
				minSubArrayLen = r - l
			}

			// scale up windows
			t := nums[l]
			l++
			sum -= t

			//fmt.Printf("--> l:%v,r:%v,result:%v,sum:%v\n", l, r, minSubArrayLen, sum)
		}
	}

	if minSubArrayLen == len(nums)+1 {
		return 0
	}

	return minSubArrayLen

}
