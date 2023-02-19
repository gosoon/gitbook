package main

type stack struct {
	topIndex int
	nums     []int
}

func NewStack() *stack {
	return &stack{topIndex: -1}
}
func (s *stack) Pop() int {
	if s.topIndex == -1 {
		return 0
	}
	num := s.nums[s.topIndex]
	s.topIndex--
	s.nums = s.nums[:s.topIndex+1]
	return num
}

func (s *stack) Push(num int) {
	s.nums = append(s.nums, num)
	s.topIndex++
}

func (s *stack) Peek() int {
	if s.topIndex == -1 {
		return 0
	}
	return s.nums[s.topIndex]
}

func (s *stack) IsEmpty() bool {
	return s.topIndex == -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaterNums := generateNextGreaterNums(nums2)
	nums2Map := make(map[int]int)
	for idx, num := range nums2 {
		nums2Map[num] = idx
	}

	var res []int
	for _, num := range nums1 {
		if idx, found := nums2Map[num]; found {
			res = append(res, nextGreaterNums[idx])
		}
	}
	return res
}

func generateNextGreaterNums(nums []int) []int {
	s := NewStack()
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for !s.IsEmpty() && s.Peek() <= nums[i] {
			s.Pop()
		}

		res[i] = s.Peek()
		if s.IsEmpty() {
			res[i] = -1
		}
		s.Push(nums[i])
	}
	return res
}
