/**
In Go's `regexp` package, you can use anchoring and flags to modify how a regular expression (regex) pattern matches text. Anchoring specifies where the pattern should start and end matching, and flags provide additional behavior modifications. Let's explore how anchoring and flags work:

### Anchoring:

1. **Start Anchoring (^)**:
   - `^` (caret) is used to anchor the pattern at the beginning of the input text.
   - When `^` is placed at the beginning of a regex pattern, it ensures that the pattern only matches text that starts with the specified pattern.
   - For example, the pattern `^Go` matches strings that start with "Go," like "GoLang" or "Gopher."

2. **End Anchoring ($)**:
   - `$` (dollar sign) is used to anchor the pattern at the end of the input text.
   - When `$` is placed at the end of a regex pattern, it ensures that the pattern only matches text that ends with the specified pattern.
   - For example, the pattern `end$` matches strings that end with "end," like "friend" or "legend."

3. **Both Start and End Anchoring (^ and $)**:
   - Combining `^` and `$` in a regex pattern ensures that the entire input text matches the pattern exactly.
   - For example, the pattern `^Go$` matches the exact string "Go" and nothing else.

### Flags:

1. **Case-Insensitive Matching (i)**:
   - The `i` flag makes the regex pattern case-insensitive, allowing it to match both uppercase and lowercase letters.
   - For example, the pattern `(?i)go` with the flag matches "go," "Go," "gO," and "G0."

2. **Global Matching (g)**:
   - The `g` flag is not a standard Go regex flag but is used in some other regex implementations.
   - In Go's `regexp` package, you typically use `FindAllString` to achieve global matching. This function returns all matches in the input text as a slice.

3. **Multiline Matching (m)**:
   - The `m` flag is used to enable multiline matching mode.
   - In multiline mode, `^` and `$` match the start and end of each line in the input text, rather than just the start and end of the entire text.
   - This is useful when processing text with multiple lines.

4. **Dot-All Mode (s)**:
   - The `s` flag is used to make the dot `.` in the pattern match any character, including newline characters `\n`.
   - By default, the dot `.` matches any character except newline. Enabling `s` allows you to match across multiple lines.

Here's an example demonstrating how anchoring and flags work:

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Go is a great programming language.\nGolang is fun to learn."
	pattern := regexp.MustCompile(`^Go`) // Anchored at the start
	match := pattern.FindString(text)
	fmt.Println("Start Anchoring:", match)

	pattern = regexp.MustCompile(`language.$`) // Anchored at the end
	match = pattern.FindString(text)
	fmt.Println("End Anchoring:", match)

	pattern = regexp.MustCompile(`(?i)go`) // Case-insensitive
	matches := pattern.FindAllString(text, -1)
	fmt.Println("Case-Insensitive:", matches)

	pattern = regexp.MustCompile(`(?m)^Go`) // Multiline
	matches = pattern.FindAllString(text, -1)
	fmt.Println("Multiline Anchoring:", matches)

	pattern = regexp.MustCompile(`(?s).$`) // Dot-All
	match = pattern.FindString(text)
	fmt.Println("Dot-All Mode:", match)
}
```

In this example, we use various anchoring and flags to match patterns within the input text, demonstrating how they modify the behavior of the regex pattern matching.**/

package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Go is a great programming language.\nGolang is fun to learn."
	pattern := regexp.MustCompile(`^Go`) // Anchored at the start
	match := pattern.FindString(text)
	fmt.Println("Start Anchoring:", match)

	pattern = regexp.MustCompile(`language.$`) // Anchored at the end
	match = pattern.FindString(text)
	fmt.Println("End Anchoring:", match)

	pattern = regexp.MustCompile(`(?i)go`) // Case-insensitive
	matches := pattern.FindAllString(text, -1)
	fmt.Println("Case-Insensitive:", matches)

	pattern = regexp.MustCompile(`(?m)^Go`) // Multiline
	matches = pattern.FindAllString(text, -1)
	fmt.Println("Multiline Anchoring:", matches)

	pattern = regexp.MustCompile(`(?s).$`) // Dot-All
	match = pattern.FindString(text)
	fmt.Println("Dot-All Mode:", match)
}
