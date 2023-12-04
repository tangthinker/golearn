package main

import (
	"fmt"
	"regexp"
)

func main() {
	phoneNumberRegexp := "^1(3\\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$"
	phoneNumber := "18613231006"
	matched, err := regexp.Match(phoneNumberRegexp, []byte(phoneNumber))
	if err == nil {
		fmt.Println(matched)
	}
	matched, err = regexp.MatchString(phoneNumberRegexp, phoneNumber)
	if err == nil {
		fmt.Println(matched)
	}
	reg, _ := regexp.Compile(phoneNumberRegexp)
	replaced := reg.ReplaceAllString(phoneNumber, "1##########")
	replacedFunc := reg.ReplaceAllStringFunc(phoneNumber, func(s string) string {
		b := []byte(s)
		for i := 3; i < 7; i++ {
			b[i] = '#'
		}
		return string(b)
	})
	fmt.Println(replaced)
	fmt.Println(replacedFunc)
}
