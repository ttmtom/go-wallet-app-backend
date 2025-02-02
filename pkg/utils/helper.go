package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Validate(input string, regex string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}

func GetUserInput(p *string) {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	input, err := buf.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	} else {
		input = input[:len(input)-1]
		*p = input
	}
}
