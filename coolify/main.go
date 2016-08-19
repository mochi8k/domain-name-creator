package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const vowelString = "aiueoAiueo"

func isVowel(s string) bool {
	return strings.Contains(vowelString, s)
}

func randBool() bool {
	return rand.Intn(2) == 0
}

func duplicateVowel(word []byte, i int) []byte {
	if i < 0 {
		return word
	}
	return append(word[:i+1], word[i:]...)
}

func removeVowel(word []byte, i int) []byte {
	if i < 0 {
		return word
	}
	return append(word[:i], word[i+1:]...)
}

func transform(word []byte) []byte {
	var vI int = -1
	for i, char := range word {
		if !isVowel(string(char)) {
			continue
		}

		if randBool() {
			vI = i
		}

		if randBool() {
			word = duplicateVowel(word, vI)
		} else {
			word = removeVowel(word, vI)
		}

	}
	return word
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())

		if randBool() {
			word = transform(word)
		}

		fmt.Println(string(word))
	}
}
