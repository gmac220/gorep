package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var scanner *bufio.Scanner
	expression, _ := regexp.Compile(os.Args[1])

	if len(os.Args) == 1 {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, _ := os.Open(os.Args[2])
		scanner = bufio.NewScanner(file)
	}

	for scanner.Scan() {
		if expression.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}
}
