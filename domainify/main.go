package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func spaceReplacer(r rune) rune {
	if unicode.IsSpace(r) {
		return '-'
	}
	return r
}

func main() {

	if len(os.Args) > 1 {
		tlds = strings.Split(os.Args[1], ",")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		text := strings.ToLower(s.Text())
		var newText []rune

		for _, r := range text {
			v := spaceReplacer(r)
			if strings.ContainsRune(allowedChars, v) {
				newText = append(newText, v)
			}
		}

		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
