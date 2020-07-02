package main

import "fmt"

func piTable(val string) []int {
	// Step 1: Base piTable
	pi := make([]int, len(val)+1)

	// Step 2: curMatch
	cM := 0

	for i := 1; i < len(val); i++ {
		// When there is a match
		if val[i] == val[cM] {
			pi[i+1] = cM + 1
			cM++
			continue
		}

		// Base case
		cM = 0
		pi[i+1] = 0
	}

	// Step 4: Return
	return pi
}

func kmp(str string, pat string) {
	pi := piTable(pat)

	curPos := 0
	for i := 0; i < len(str); i++ {
		// If equal
		if str[i] == pat[curPos] {
			curPos++
			if curPos == len(pat) {
				fmt.Println("Pattern found at:", i+1-curPos)
				curPos = pi[curPos]
				i--
			}
			continue
		}

		// If not equal
		if curPos != 0 {
			curPos = pi[curPos]
			i--
		}
	}
}

func main() {
	kmp("AABAACAADAABAABA", "AABA")
}
