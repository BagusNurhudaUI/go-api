package static

import (
	"fmt"
	"strings"
)

func Init() {
	fmt.Println("this is static")

	// strings.Contains
	str := "john wick is really good man"
	checkStr := "yaya"
	var isExists = strings.Contains(str, checkStr)
	fmt.Printf("check strings.Contains have %s on %s is %v \n", checkStr, str, isExists)

	var string1 = strings.Split("the dark    knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "    knight"]
	fmt.Println("len string1", len(string1))
	fmt.Println(string1[1])

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	string1 = strings.Fields("the dark    knight")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

}

func reverseVowels(s string) string {
	vowel := "aiueoAIUEO"
	lo, hi := 0, len(s)-1
	runes := make([]rune, len(s))
	for i, r := range s {
		runes[i] = r
	}

	for lo < hi {
		fmt.Println(lo, hi)
		if !strings.Contains(vowel, string(runes[hi])) {
			hi--
			continue
		}

		if !strings.Contains(vowel, string(runes[lo])) {
			lo++
			continue
		}

		runes[lo], runes[hi] = runes[hi], runes[lo]
		lo++
		hi--
	}
	fmt.Println(string(runes))
	return string(runes)
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	if len(str1)%len(str2) == 0 {
		return str2
	}

	gcdLength := gcd(len(str1), len(str2))
	result := str1[:gcdLength]

	return result
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
		fmt.Println(a, b)
	}
	return a
}
