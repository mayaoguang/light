package algorithm

// IsSubsequence 判断子序列 了code 392
func IsSubsequence(s, t string) bool {
	if len(s) > len(t) {
		return false
	}
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// 进阶的功能 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
var allSubStr = make(map[string]struct{})

func SubSets(myStr, subset string, start int, allSubSeq map[string]struct{}) {
	allSubSeq[subset] = struct{}{}
	for i := start; i < len(myStr); i++ {
		SubSets(myStr, subset+string(myStr[i]), i+1, allSubSeq)
	}
}

func AllSubSequence(t string) {
	SubSets(t, "", 0, allSubStr)
}

func IsSubSeq(s string) bool {
	_, ok := allSubStr[s]
	return ok
}

func isSubsequence(s string, t string) bool {
	AllSubSequence(t)
	_, ok := allSubStr[s]
	return ok
}
