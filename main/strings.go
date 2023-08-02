package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func stringsStart() {
	fmt.Println(nextASCII(byte(100)))
	fmt.Println(prevASCII(byte(99)))
	fmt.Println(shiftASCII("abc1234", -511))
	fmt.Println(isASCII("sad"))
	fmt.Println(isASCII("хай"))
	fmt.Println(latinLetters("John Уолтер"))
	fmt.Println(latinLetters("11 ! t e    s t %"))
	fmt.Println(generateSelfStory("Vlad", 25, 10.00000025))
}

func nextASCII(b byte) byte {
	return b + uint8(1)
}

func prevASCII(b byte) byte {
	return b - uint8(1)
}

func shiftASCII(s string, step int) string {
	bytesRes := []byte{}
	for i := 0; i < len(s); i++ {
		bytesRes = append(bytesRes, byte(s[i])+uint8(step))
	}

	return string(bytesRes)
}

func isASCII(s string) bool {
	length := utf8.RuneCountInString(s)
	return len(s) == length
}

func latinLetters(s string) string {
	resArr := make([]rune, 0)
	for _, rn := range s {
		isLatin := unicode.Is(unicode.Latin, rune(rn))
		if isLatin {
			resArr = append(resArr, rn)
		}
	}
	return string(resArr)
}

func generateSelfStory(name string, age int, money float64) string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("Hello! My name is %s. ", name))
	builder.WriteString(fmt.Sprintf("I'm %d y.o. ", age))
	builder.WriteString(fmt.Sprintf("And I also have $%.2f in my wallet right now.", money))
	return builder.String()
}
