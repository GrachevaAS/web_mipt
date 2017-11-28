package main

import (
	"unicode"
)

func RemoveEven(slice []int)(noeven [] int) {
	for _, num := range(slice) {
		if num % 2 == 1 {
			noeven = append(noeven, num)
		}
	}
	return
}

func PowerGenerator(num int) (func() int) {
  arg := num
  ans := 1
  return func() (int) {
    ans *= arg
    return ans
  }
}

func DifferentWordsCount(str string)(int) {
	words := make(map[string]bool)
	current_word := ""
	for _, symb := range(str) {
		if unicode.IsLetter(symb) {
			current_word = current_word + string(unicode.ToLower(rune(symb)))
		} else {
			if (len(current_word) != 0) {
				words[current_word] = true
			}
			current_word = ""
		}
	}
	if (len(current_word) != 0) {
		words[current_word] = true
	}
	return len(words)
}