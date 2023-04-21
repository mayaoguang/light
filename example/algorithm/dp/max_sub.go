package dp

// LengthOfASC 无序数组求最长上升子序列
func LengthOfASC(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make([]int, len(nums))
	result := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// LongestCommonSubsequence 最长公共子序列长度
func LongestCommonSubsequence(text1 string, text2 string) int {
	h := len(text1) + 1
	w := len(text2) + 1
	m := make([][]int, h)
	for i := 0; i < h; i++ {
		m[i] = make([]int, w)
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if text1[i-1] == text2[j-1] {
				m[i][j] = m[i-1][j-1] + 1
			} else {
				if m[i-1][j] > m[i][j-1] {
					m[i][j] = m[i-1][j]
				} else {
					m[i][j] = m[i][j-1]
				}
			}
		}
	}
	return m[h-1][w-1]
}
