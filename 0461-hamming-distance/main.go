package main

func hammingDistance(x int, y int) int {
	tmp := x ^ y
	result := 0
	for i := 0; i < 32; i++ {
		if tmp&1 > 0 {
			result++
		}
		tmp >>= 1
	}
	return result
}
