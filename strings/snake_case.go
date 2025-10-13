package strings

import stdstrings "strings"

// ToSnakeCase converts a concatenated string to snake_case by segmenting it into
// dictionary words and joining them with underscores in lowercase.
//
// Examples: "helloworldfromme" → "hello_world_from_me", "opensource" → "open_source"
//
// Falls back to returning the input in lowercase if no segmentation is found.
func ToSnakeCase(input string) string {
	ensureInitialized()
	words := segment(stdstrings.ToLower(input))
	if words == nil {
		return stdstrings.ToLower(input)
	}
	return stdstrings.Join(words, "_")
}
