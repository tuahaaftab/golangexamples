package golangexamples

import "github.com/ehteshamz/greetings"

//EZGreetings prints greetings
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

//ConcatSlice concatinates slice
func ConcatSlice(sliceToConcat []byte) string {
	s := ""
	for i, v := range sliceToConcat {
		if i == 0 {
			s = string(v)
		} else {
			s = s + "-" + string(v)
		}
	}

	return s
}

//Encrypts slice with ceaser cypher
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i, _ := range sliceToEncrypt {
		sliceToEncrypt[i] += byte(ceaserCount)
	}
}
