package main

type NumArray struct {
	Nums   []int
	PreSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{Nums: nums, PreSum: preSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.PreSum[right+1] - this.PreSum[left]
}
