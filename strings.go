// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var result string = string_reverser("hello world")
	fmt.Println(result)
}

func string_reverser(our_string string) string {
	var length int = len(our_string)
	var our_string_array = strings.Split(our_string, "")
	reversed_string := make([]string, length)
	j := 0
	for i := length-1; i >= 0; i-- {
		reversed_string[j] = our_string_array[i]
		j++
	}
	return strings.Join(reversed_string,"")
}