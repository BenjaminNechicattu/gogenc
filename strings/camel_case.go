package strings

import stdstrings "strings"

// ToCamelCase converts a concatenated string to camelCase by segmenting it into
// dictionary words, keeping the first word lowercase and capitalizing the rest.
//
// Examples: "helloworldfromme" → "helloWorldFromMe", "opensource" → "openSource"
//
// Falls back to returning the input in lowercase if no segmentation is found.
func ToCamelCase(input string) string {
	ensureInitialized()
	words := segment(stdstrings.ToLower(input))
	if words == nil {
		return stdstrings.ToLower(input)
	}
	var result stdstrings.Builder
	for i, word := range words {
		if i == 0 {
			result.WriteString(word)
		} else {
			result.WriteString(capitalize(word))
		}
	}
	return result.String()
}
