package main

import (
	"fmt"
	"regexp"
)

// \d+: Matches one or more digits.
// [A-Za-z]+: Matches one or more uppercase or lowercase letters.
// ^[0-9]{5}$: Matches a string containing exactly five digits.
// \bword\b: Matches the word "word" as a whole word.

func main() {
	// 	Creating a Regular Expression:
	// You create a regular expression pattern using the regexp.Compile() or regexp.MustCompile() function.
	// The former returns an error if the pattern is invalid,
	// Using regexp.MustCompile
	pattern := regexp.MustCompile(`\d+`) //matiches one or more digits

	// 	Matching a String:
	// You can use methods like MatchString(), FindString(), and FindStringSubmatch() to perform various types of matches.
	text := "There are 42 apples and 123 oranges."
	matched := pattern.MatchString(text)
	fmt.Println(matched)

	//find string
	result := pattern.FindString(text)
	fmt.Println(result)

	//Use FindAllString() to find all occurrences of the pattern in a string.
	allMatches := pattern.FindAllString(text, -1) // -1 means find all occurrences
	fmt.Println("find all string", allMatches)

	//You can replace matches using ReplaceAllString().
	newText := pattern.ReplaceAllString(text, "X")
	fmt.Println("replaced string", newText)

	//Parentheses () in the pattern define capturing groups. You can extract submatches from capturing groups.
	// 	Suppose you have the following regex pattern: (\d+)\s+(\w+).
	// (\d+) is the first capturing group, which matches one or more digits.
	// \s+ matches one or more whitespace characters.
	// (\w+) is the second capturing group, which matches one or more word characters (letters, digits, or underscores).
	pattern2 := regexp.MustCompile(`(\d+)\s+(\w+)`)
	submatches := pattern2.FindStringSubmatch("42 apples")
	fmt.Println("submatch", submatches)

	//Always check for errors when working with regular expressions, especially if the pattern is dynamic.
	pattern, err := regexp.Compile(`[a-z+`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
	}

}
