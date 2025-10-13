package strings

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	stdstrings "strings"
)

var wordDict map[string]bool
var memo map[string][]string
var initialized bool

func ensureInitialized() {
	if !initialized {
		wordDict = make(map[string]bool)
		memo = make(map[string][]string)
		loadDictionary()
		initialized = true
	}
}
func loadDictionary() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("runtime.Caller failed")
	}
	dir := filepath.Dir(filename)
	filePath := filepath.Join(dir, "20k.txt")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := stdstrings.TrimSpace(scanner.Text())
		if word != "" {
			wordDict[word] = true
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func segment(input string) []string {
	if result, exists := memo[input]; exists {
		return result
	}

	// Try to find the best segmentation by trying all possible first words
	// and picking the one that leads to a valid complete segmentation
	var bestResult []string
	var bestScore int // Prefer solutions with fewer, longer words

	for i := 1; i <= len(input); i++ {
		prefix := input[:i]
		if wordDict[prefix] {
			suffix := input[i:]
			if suffix == "" {
				// Complete match found
				result := []string{prefix}
				memo[input] = result
				return result
			}

			// Try to segment the rest
			rest := segment(suffix)
			if rest != nil {
				result := append([]string{prefix}, rest...)
				// Calculate score: prefer fewer words with longer average length
				score := calculateScore(result)
				if bestResult == nil || score > bestScore {
					bestResult = result
					bestScore = score
				}
			}
		}
	}

	memo[input] = bestResult
	return bestResult
}

// calculateScore gives higher scores to segmentations with fewer, longer words
func calculateScore(words []string) int {
	if len(words) == 0 {
		return 0
	}

	totalLength := 0
	for _, word := range words {
		totalLength += len(word)
	}

	// Score = average word length * 100 - number of words
	// This prefers longer words and fewer segments
	avgLength := totalLength / len(words)
	return avgLength*100 - len(words)
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return stdstrings.ToUpper(word[:1]) + stdstrings.ToLower(word[1:])
}

// ToPascalCase converts a concatenated string to PascalCase by segmenting it into
// dictionary words and capitalizing each word.
//
// Examples: "helloworldfromme" → "HelloWorldFromMe", "opensource" → "OpenSource"
//
// Falls back to capitalizing the input if no segmentation is found.
func ToPascalCase(input string) string {
	ensureInitialized()
	words := segment(stdstrings.ToLower(input))
	if words == nil {
		return capitalize(input)
	}
	var result stdstrings.Builder
	for _, word := range words {
		result.WriteString(capitalize(word))
	}
	return result.String()
}
