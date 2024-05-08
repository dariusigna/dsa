package main

import "fmt"

var phoneMap map[string]string

func generateCombination(s string, targetLen int, solution []byte, results *[]string) {
	if len(solution) == targetLen {
		*results = append(*results, string(solution))
		return
	}

	if len(s) == 0 {
		return
	}

	chars, _ := phoneMap[string(s[0])]
	for _, c := range chars {
		solution = append(solution, byte(c))
		generateCombination(s[1:], targetLen, solution, results)
		solution = solution[:len(solution)-1]
	}
}

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	phoneMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var results []string
	generateCombination(digits, len(digits), []byte{}, &results)
	return results
}

func main() {
	fmt.Println(letterCombinations("23"))
}
