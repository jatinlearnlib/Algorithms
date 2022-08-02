package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	str := "ABCDE"
	rev, err := reverseString1(str)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(rev)
}

/*
	1. String is immutable so we can't change a particular index value of string variable or type.
	2. One way to alter string type is to convert string into slice of rune type and then change the particular
	   byte value and then convert slice of rune back to string. Rune is an alias of int32 type that we should
	   use when we are dealing with worldwide languages.
	3. Second way is similar to previous one but instead of using slice of rune we can use slice of byte if we
	   are dealing with ASCII characters or single byte characters only. As byte is an alias on uint8 and takes
	   less memory as compared to rune that takes 4 byte.
*/

//Approach 1
func reverseString1(str string) (string, error) {
	x := []byte(str)
	if len(str) <= 0 {
		return "", fmt.Errorf("String is empty")
	}

	for i, j := 0, len(x)-1; i < len(x)/2; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	return string(x), nil

}

//Approach 2
func reverseString2(str string) (string, error) {
	x := make([]string, 0)

	if len(str) <= 0 {
		return "", fmt.Errorf("String is empty")
	}

	for i := len(str) - 1; i >= 0; i-- {
		x = append(x, string(str[i]))
	}

	return strings.Join(x, ""), nil
}
