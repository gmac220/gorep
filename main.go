package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	expression, _ := regexp.Compile(os.Args[1])

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if expression.MatchString(scanner.Text()) {
			fmt.Println(scanner.Text())
		}
	}
}
