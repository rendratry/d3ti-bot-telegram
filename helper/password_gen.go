package helper

import (
	"regexp"
	"strings"
)

func Reverse(str string) string {

	//Fungsi Membalikkan String
	chars := []rune(str)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	return string(result)
}

func GeneratePassword(str string) string {
	//Replace
	str = Reverse(str)

	replacer := strings.NewReplacer(
		"a", "E", "e", "I", "i", "O", "o", "U", "u", "A",
		"A", "e", "E", "i", "I", "o", "O", "u", "U", "a",
		" ", "")

	replacedstr := replacer.Replace(str)

	letters := strings.Split(replacedstr, "")

	swapCase := func() []string {
		var newArray []string
		for _, element := range letters {
			//rExp := regexp.MustCompile(`[^aeiouAEIOU]`)
			rUpper := regexp.MustCompile(`[b-df-hj-np-tv-z]`)
			rLower := regexp.MustCompile(`[B-DF-HJ-NP-TV-Z]`)

			if rUpper.MatchString(element) {
				newArray = append(newArray, strings.ToUpper(element))
			} else if rLower.MatchString(element) {
				newArray = append(newArray, strings.ToLower(element))
			} else {
				newArray = append(newArray, element)
			}
		}
		return newArray
	}

	return strings.Join(swapCase(), "")
}
