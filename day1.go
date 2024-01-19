package main

import (
	"fmt"
	"strings"
)

type EdgeLengthCache struct {
	cache     map[string]bool
	perimeter int
}

func (r *EdgeLengthCache) SetEdgeLengthCache(row, col int) {
	key := fmt.Sprintf("%d_%d", row, col)
	r.cache[key] = true

	// 查看左边和上边是否有接壤，如果接壤要减去2
	if row != 0 {
		_, found := r.cache[fmt.Sprintf("%d_%d", row-1, col)]
		if found {
			r.perimeter -= 2
		} else {
			r.cache[key] = true
		}
	}
	if col != 0 {
		_, found := r.cache[fmt.Sprintf("%d_%d", row, col-1)]
		if found {
			r.perimeter -= 2
		} else {
			r.cache[key] = true
		}
	}
}

func islandPerimeter(grid [][]int) int {
	cache := &EdgeLengthCache{}
	cache.cache = make(map[string]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 1 {
				cache.perimeter += 4
				cache.SetEdgeLengthCache(row, col)
			}
		}
	}

	return cache.perimeter
}

func findMaxConsecutiveOnes(nums []int) int {
	maxConsecutiveOnes := 0
	start := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if start == -1 {
				start = i
				end = i
			} else {
				end++
			}
		} else {
			if start != -1 {
				maxConsecutiveOnes = max(maxConsecutiveOnes, end-start+1)
				start = -1
				end = -1
			}
		}
	}

	if start != -1 {
		maxConsecutiveOnes = max(maxConsecutiveOnes, end-start+1)
	}

	return maxConsecutiveOnes
}

// 输入：timeSeries = [1,4], duration = 2
// 输出：4
// 解释：提莫攻击对艾希的影响如下：
// - 第 1 秒，提莫攻击艾希并使其立即中毒。中毒状态会维持 2 秒，即第 1 秒和第 2 秒。
// - 第 4 秒，提莫再次攻击艾希，艾希中毒状态又持续 2 秒，即第 4 秒和第 5 秒。
// 艾希在第 1、2、4、5 秒处于中毒状态，所以总中毒秒数是 4 。
// 示例 2：

// 输入：timeSeries = [1,2], duration = 2
// 输出：3
// 解释：提莫攻击对艾希的影响如下：
// - 第 1 秒，提莫攻击艾希并使其立即中毒。中毒状态会维持 2 秒，即第 1 秒和第 2 秒。
// - 第 2 秒，提莫再次攻击艾希，并重置中毒计时器，艾希中毒状态需要持续 2 秒，即第 2 秒和第 3 秒。
// 艾希在第 1、2、3 秒处于中毒状态，所以总中毒秒数是 3 。

func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	start := timeSeries[0]
	end := timeSeries[0] + duration

	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] < end {
			end = timeSeries[i] + duration
		} else {
			ans += end - start
			start = timeSeries[i]
			end = timeSeries[i] + duration
		}
	}
	ans += end - start
	return ans
}

func findPoisonedDuration1(timeSeries []int, duration int) int {
	res := 0
	for i := 0; i < len(timeSeries); i++ {
		if i == 0 {
			res += duration
		} else {
			lastPoisonedEnd := timeSeries[i-1] + duration
			if lastPoisonedEnd > timeSeries[i] {
				res += duration - lastPoisonedEnd + timeSeries[i]
			} else {
				res += duration
			}
		}
	}
	return res
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	monotonicStack := make([]int, len(nums2))
	monotonicStackLen := 0
	nextGreaterElementMap := make(map[int]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		if monotonicStackLen == 0 {
			nextGreaterElementMap[nums2[i]] = -1

			monotonicStack[monotonicStackLen] = nums2[i]
			monotonicStackLen++

			continue
		}

		if monotonicStack[monotonicStackLen-1] > nums2[i] {
			// 栈顶大
			fmt.Printf("monotonicStack[monotonicStackLen-1] %+v\n", monotonicStack[monotonicStackLen-1])
			nextGreaterElementMap[nums2[i]] = monotonicStack[monotonicStackLen-1]

			monotonicStack[monotonicStackLen] = nums2[i]
			monotonicStackLen++
		} else {
			// 栈顶小，循环吐出元素
			for j := monotonicStackLen - 1; j >= 0; j-- {
				if monotonicStack[monotonicStackLen-1] < nums2[i] {
					monotonicStackLen--
				} else {
					break
				}
			}
			if monotonicStackLen == 0 {
				nextGreaterElementMap[nums2[i]] = -1

				monotonicStack[monotonicStackLen] = nums2[i]
				monotonicStackLen++
			} else {
				nextGreaterElementMap[nums2[i]] = monotonicStack[monotonicStackLen-1]
				monotonicStack[monotonicStackLen] = nums2[i]
				monotonicStackLen++
			}

		}

		fmt.Printf("monotonicStack： %+v\n", monotonicStack)

	}

	res := make([]int, 0)
	for _, val := range nums1 {
		res = append(res, nextGreaterElementMap[val])
	}
	return res
}

func findWords(words []string) []string {
	validMap := make(map[rune]int)
	for _, val := range "qwertyuiop" {
		validMap[val] = 1
	}
	for _, val := range "asdfghjkl" {
		validMap[val] = 2
	}
	for _, val := range "zxcvbnm" {
		validMap[val] = 3
	}

	ans := make([]string, 0)
	for _, word := range words {
		label := 0
		for _, val := range strings.ToLower(word) {
			if label == 0 {
				label = validMap[val]
				continue
			}

			if label != validMap[val] {
				label = -1
				break
			}
		}
		if label != -1 {
			ans = append(ans, word)
		}

	}

	return ans
}

func main() {
	// grid := [][]int{
	//  {0, 1, 0, 0},
	//  {1, 1, 1, 0},
	//  {0, 1, 0, 0},
	//  {1, 1, 0, 0},
	// }
	// fmt.Println(islandPerimeter(grid))

	// nums := []int{1, 1, 0, 1, 1, 1}
	// fmt.Println(findMaxConsecutiveOnes(nums))

	// timeSeries := []int{1, 4}
	// fmt.Println(findPoisonedDuration(timeSeries, 2))

	// nums1 := []int{4, 1, 2}
	// nums2 := []int{1, 3, 4, 2}
	// fmt.Printf("%+v\n", nextGreaterElement(nums1, nums2))

	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Printf("%+v\n", findWords(words))

	return
}
