package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverse(t *testing.T) {
	t.Run("Test Reverse", func(t *testing.T) {
		testString := "benjamin"
		expected := "nimajneb"
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse", func(t *testing.T) {
		testString := "benjaminNechicattu"
		expected := "nimajneb"
		resp := Reverse(testString)
		require.NotEqual(t, expected, resp)
	})
	t.Run("Test Reverse with empty string", func(t *testing.T) {
		testString := ""
		expected := ""
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse with single character", func(t *testing.T) {
		testString := "a"
		expected := "a"
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse with special characters", func(t *testing.T) {
		testString := "!@#$%^&*()"
		expected := ")(*&^%$#@!"
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse with spaces", func(t *testing.T) {
		testString := "hello world"
		expected := "dlrow olleh"
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse with numbers", func(t *testing.T) {
		testString := "1234567890"
		expected := "0987654321"
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse with mixed characters", func(t *testing.T) {
		testString := "abc123!@#"
		expected := "#@!321cba"
		resp := Reverse(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test Reverse with long string", func(t *testing.T) {
		testString := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
		expected := ".auqila angol erod et erobal tu tiddnI morpmet sdoeS .tile gnidicniap tsetnoc ,tema tis rolod muspi meroL"
		resp := Reverse(testString)
		require.NotEqual(t, expected, resp)
	})
}
