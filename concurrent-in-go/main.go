package main

func main() {}

func moveZeroToTain(nums []int) {
	// left: loop index
	// right: determine non-zero position
	left, right := 0, 0

	for right < len(nums) {
		for nums[right] == 0 {
			right++
		}
		nums[left] = nums[right]
		left++
		right++
	}

	for left < len(nums) {
		nums[left] = 0
		left++
	}
}
