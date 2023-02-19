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

// 73, 74, 75, 71, 69, 72, 76, 73
func dailyTemperatures(temperatures []int) []int {
	return generateNextDailyTemperatures(temperatures)
}

func generateNextDailyTemperatures(nums []int) []int {
	s := NewStack()
	l := len(nums)
	res := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		for !s.IsEmpty() && nums[s.Peek()] <= nums[i] {
			s.Pop()
		}

		res[i] = s.Peek() - i
		if s.IsEmpty() {
			res[i] = 0
		}
		s.Push(i)
	}
	return res
}
