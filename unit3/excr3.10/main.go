package main

import (
	"bytes"
	"fmt"
)

func insertComma(s string) string {
	var buff bytes.Buffer
	if len(s) <= 3 {
		return s
	} else if len(s)%3 == 0 {
		for i, rune := range s {
			buff.WriteRune(rune)
			if (i+1)%3 == 0 && i != len(s)-1 {
				buff.WriteRune(',')
			}
		}
	} else {
		initialDigits := len(s) % 3
		initialString := s
		for i := 0; i <= len(initialString)-1; i++ {
			buff.WriteByte(initialString[i])
			if i == initialDigits-1 {
				buff.WriteRune(',')
				newString := initialString[i+1:]
				for i := 0; i <= len(newString)-1; i++ {
					buff.WriteByte(newString[i])
					if (i+1)%3 == 0 && i != len(newString)-1 {
						buff.WriteRune(',')
					}
				}
				return buff.String()
			}
		}
	}
	return buff.String()
}

func main() {
	fmt.Println(insertComma("1234567890121"))
}
