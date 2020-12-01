// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	// String reverser
	fmt.Println("String reverserver:")
	result := string_reverser("hello world")
	fmt.Println(result)

	// Anagram checker
	fmt.Println("Anagram checker")
	if anagram_checker("tenet", "tenet1") == true {
		fmt.Println("It's anagram")
	} else {
		fmt.Println("It's not anagram")
	}

	// Word flipper
	fmt.Println("Word flipper")
	fmt.Println(word_flipper("hello how are you?"))
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

func sort_string(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
}

func anagram_checker(str1 string, str2 string) bool {
	str1 = sort_string(str1)
	str2 = sort_string(str2)

	if len(str1) != len(str2) {
		return false
	}

	str1_array := strings.Split(str1, "")
	str2_array := strings.Split(str2, "")

	for i := 0; i < len(str1); i++ {
		if str1_array[i] != str2_array[i] {
			return false
		}
	}
	
	return true
}

func word_flipper(s string) string {
	words := strings.Split(s, " ")
	length := len(words)
	reversed_words := make([]string, length)
	j := 0

	for i := length - 1; i >= 0; i-- {		
		reversed_words[j] = words[i]
		j++
	}

	return strings.Join(reversed_words, " ")
}