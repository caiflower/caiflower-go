package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i, _ := range dp {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len1][len2]
}

func breakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}

	l := len(palindrome) / 2
	//if len(palindrome)&1 != 0 {
	//	l += 1
	//}

	for i := 0; i < l; i++ {
		if palindrome[i] != 'a' {
			return palindrome[:i] + "a" + palindrome[i+1:]
		}
	}
	return palindrome[:len(palindrome)-1] + "b"
}

func testSlice() {
	arr := [10]int{}
	s := arr[0:5]
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	fmt.Println(arr)
	fmt.Println("--------")
	s = append(s, 6)
	fmt.Println(s)
	fmt.Println(arr)
	fmt.Println(cap(s))
}

func main() {
	//fmt.Println(longestCommonSubsequence("abcde", "ace"))

	testSlice()
}
