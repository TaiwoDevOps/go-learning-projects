package main

import "strings"

func AreAlmostEqual(s1 string, s2 string) bool {
	var tempS1 []string //split the first string into an array
	var temp []string
	var tempIndexes []int
	var charLen int
	var swappedString string

	// check if the strings are the same
	if s1 == s2 {
		return true
	}

	// check if each character at the giving index is the same
	// "apple", "elppa"
	for i := range s1 {
		tempS1 = append(tempS1, string(s1[i]))
		if s1[i] != s2[i] {
			// check if the length of mismatched chars is less than 2
			if charLen < 2 {
				// if it is, assign the the first two mismatched chars and their indexes to an array
				// Note: for optimization, this is where the hashmap  might be used
				temp = append(temp, string(s1[i]))
				tempIndexes = append(tempIndexes, i)
			}
			charLen++
		}
	}

	// "apple", "elppa"
	//check to see if we have at least 2 mismatched chars in a string and we have 2 mismatched chars held in an array
	if charLen >= 2 && len(temp) == 2 {
		// swapped the mismatched chars in the array holding the string
		tempS1[tempIndexes[0]] = temp[1]
		tempS1[tempIndexes[1]] = temp[0]
	}
	// convert the array to a string
	swappedString = strings.Join(tempS1, "")
	// use this condition to return the result of the check
	return swappedString == s2
}

// Key learnings :
// 1. Array manipulation
//->	// how to insert an element into an array either at a specific index or not
//->	// how to remove an element from an array
//->	// how to convert an array to a string using join
//->	// how to check the length of an array
//->	// types of loops functions in golang
// 2. String manipulation
//->	// learnt about some string methods eg... equal, len and join
//->	// how to split a string into an array
