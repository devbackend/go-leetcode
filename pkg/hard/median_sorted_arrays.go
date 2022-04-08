package hard

import (
	"math"
)

// FindMedianSortedArrays for https://leetcode.com/problems/median-of-two-sorted-arrays/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArrays(nums2, nums1)
	}

	total := len(nums1) + len(nums2)
	half := total / 2
	isOdd := total%2 == 1

	firstRightIx := len(nums1) / 2

	for {
		secondRightIx := half - firstRightIx

		firstLeftIx := firstRightIx - 1
		secondLeftIx := secondRightIx - 1

		firstLeft, firstRight, secondLeft, secondRight := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt

		if 0 <= firstLeftIx && firstLeftIx < len(nums1) {
			firstLeft = nums1[firstLeftIx]
		}

		if 0 <= firstRightIx && firstRightIx < len(nums1) {
			firstRight = nums1[firstRightIx]
		}

		if 0 <= secondLeftIx && secondLeftIx < len(nums2) {
			secondLeft = nums2[secondLeftIx]
		}

		if 0 <= secondRightIx && secondRightIx < len(nums2) {
			secondRight = nums2[secondRightIx]
		}

		if firstLeft <= secondRight && secondLeft <= firstRight {
			if isOdd {
				return math.Min(float64(firstRight), float64(secondRight))
			}

			left := math.Max(float64(firstLeft), float64(secondLeft))
			right := math.Min(float64(firstRight), float64(secondRight))

			return (left + right) / 2
		}

		if firstLeft > secondRight {
			firstRightIx--
			continue
		}

		if secondLeft > firstRight {
			firstRightIx++
			continue
		}
	}
}
