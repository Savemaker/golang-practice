package leetcode

const MAX_INT_32 = 2147483647

func IncreasingTriplet(nums []int) bool {
	first := MAX_INT_32
	second := MAX_INT_32
	for i := 0; i < len(nums); i++ {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			return true
		}
	}
	return false
}
