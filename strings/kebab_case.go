package strings

import stdstrings "strings"

// ToKebabCase converts a concatenated string to kebab-case by segmenting it into
// dictionary words and joining them with hyphens in lowercase.
//
// Examples: "helloworldfromme" → "hello-world-from-me", "opensource" → "open-source"
//
// Falls back to returning the input in lowercase if no segmentation is found.
func ToKebabCase(input string) string {
	ensureInitialized()
	words := segment(stdstrings.ToLower(input))
	if words == nil {
		return stdstrings.ToLower(input)
	}
	return stdstrings.Join(words, "-")
}
