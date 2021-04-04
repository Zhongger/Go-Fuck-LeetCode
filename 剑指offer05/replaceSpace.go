package main

import (
	"fmt"
	"strings"
)

func main() {
	res := replaceSpace("We are happy.")
	fmt.Printf("%s", res)
}

func replaceSpace(s string) string {
	var res strings.Builder
	for i := range s {
		if s[i] == ' ' {
			res.WriteString("%20")
		} else {
			res.WriteByte(s[i])
		}
	}
	return res.String()

}
