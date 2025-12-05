package day3

import (
	"advent_of_code/utils"
)

func getMaxComb(nums []int, length int) int {
	currMax, maxIx := nums[0], 0
	for ix := 1; ix <= len(nums)-length; ix++ {
		if nums[ix] > currMax {
			currMax = nums[ix]
			maxIx = ix
		}
	}

	if length == 1 {
		return currMax
	}

	return utils.PowInt(10, length-1)*currMax + getMaxComb(nums[maxIx+1:], length-1)
}
