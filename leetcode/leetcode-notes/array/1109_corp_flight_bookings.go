package main

// The most optimal
func corpFlightBookings(bookings [][]int, n int) []int {
	difference := make([]int, n)

	for _, booking := range bookings {
		difference = DifferenceIncrement(difference, booking[0]-1, booking[1]-1, booking[2])
	}
	return DifferenceResult(difference)
}

/*
func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	for _, booking := range bookings {
		for i := booking[0]; i <= booking[1]; i++ {
			res[i-1] += booking[2]
		}
	}
	return res
}
*/
