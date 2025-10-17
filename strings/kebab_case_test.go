package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToKebabCase(t *testing.T) {
	t.Run("Test ToKebabCase with helloworldfromme", func(t *testing.T) {
		testString := "helloworldfromme"
		expected := "hello-world-from-me"
		resp := ToKebabCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToKebabCase with productionready", func(t *testing.T) {
		testString := "productionready"
		expected := "production-ready"
		resp := ToKebabCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToKebabCase with single word", func(t *testing.T) {
		testString := "hello"
		expected := "hello"
		resp := ToKebabCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToKebabCase with empty string", func(t *testing.T) {
		testString := ""
		expected := ""
		resp := ToKebabCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToKebabCase with testcase", func(t *testing.T) {
		testString := "testcase"
		expected := "test-case"
		resp := ToKebabCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToCamelCase with opensource", func(t *testing.T) {
		testString := "opensource"
		expected := "open-source"
		resp := ToKebabCase(testString)
		require.Equal(t, expected, resp)
	})
}
