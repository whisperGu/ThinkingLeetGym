package main

import "fmt"

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

func main() {
	// grid := [][]int{
	// 	{0, 1, 0, 0},
	// 	{1, 1, 1, 0},
	// 	{0, 1, 0, 0},
	// 	{1, 1, 0, 0},
	// }
	// fmt.Println(islandPerimeter(grid))

	// nums := []int{1, 1, 0, 1, 1, 1}
	// fmt.Println(findMaxConsecutiveOnes(nums))

	timeSeries := []int{1, 4}
	fmt.Println(findPoisonedDuration(timeSeries, 2))

	return
}
