package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file error :", fileName)
		os.Exit(1)
	}

	defer file.Close()

	var lines []string
	s := bufio.NewScanner(file)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func main() {
	transforms := readFile("transforms.txt")
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
