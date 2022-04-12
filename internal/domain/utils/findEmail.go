package utils

import (
	"fmt"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
var pattern = regexp.MustCompile(`Email:\\c+`)

func FindEmails(s string) []string {
	//var res []string
	fmt.Println(s)
	data := pattern.Find([]byte(s))
	fmt.Println(string(data), len(data))
	return nil
}

func isEmailValid(e string) bool {
	return emailRegex.MatchString(e)
}
