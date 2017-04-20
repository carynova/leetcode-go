package twoSum

func twoSum(nums []int, target int) []int {
	posMap := make(map[int]int)

	for i, val := range nums {
		tmp := target - val
		if pos, ok := posMap[tmp]; ok {
			return []int{pos, i}
		}
		posMap[val] = i
	}
	return nil
}
