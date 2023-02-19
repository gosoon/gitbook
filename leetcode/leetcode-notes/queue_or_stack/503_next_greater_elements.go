package main

/*
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
*/

func nextGreaterElements(nums []int) []int {
	return generateNextGreaterNums2(nums)
}

func generateNextGreaterNums2(nums []int) []int {
	s := NewStack()
	l := len(nums)
	res := make([]int, l)
	for i := l*2 - 1; i >= 0; i-- {
		for !s.IsEmpty() && s.Peek() <= nums[i%l] {
			s.Pop()
		}

		res[i%l] = s.Peek()
		if s.IsEmpty() {
			res[i%l] = -1
		}
		s.Push(nums[i%l])
	}
	return res
}
