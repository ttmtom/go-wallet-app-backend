package utils

import "regexp"

func Validate(input string, regex string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(input)
}
