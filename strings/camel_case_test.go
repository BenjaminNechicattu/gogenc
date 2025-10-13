package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToCamelCase(t *testing.T) {
	t.Run("Test ToCamelCase with helloworldfromme", func(t *testing.T) {
		testString := "helloworldfromme"
		expected := "helloWorldFromMe"
		resp := ToCamelCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToCamelCase with productionready", func(t *testing.T) {
		testString := "productionready"
		expected := "productionReady"
		resp := ToCamelCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToCamelCase with single word", func(t *testing.T) {
		testString := "hello"
		expected := "hello"
		resp := ToCamelCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToCamelCase with empty string", func(t *testing.T) {
		testString := ""
		expected := ""
		resp := ToCamelCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToCamelCase with opensource", func(t *testing.T) {
		testString := "opensource"
		expected := "openSource"
		resp := ToCamelCase(testString)
		require.Equal(t, expected, resp)
	})
}
