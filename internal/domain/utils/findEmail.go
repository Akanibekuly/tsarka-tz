package utils

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var pattern = regexp.MustCompile(`Email:\s+[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)

func FindEmails(s string) []string {
	var res []string
	data := pattern.FindAllString(s, -1)
	fmt.Println(data)
	for i := range data {
		fmt.Println(emailRegex.FindAllString(data[i], -1))
		res = append(res, emailRegex.FindAllString(data[i], -1)...)
	}

	return res
}

func isEmailValid(e string) bool {
	return emailRegex.MatchString(e)
}
