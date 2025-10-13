package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToPascalCase(t *testing.T) {
	t.Run("Test ToPascalCase with helloworldfromme", func(t *testing.T) {
		testString := "helloworldfromme"
		expected := "HelloWorldFromMe"
		resp := ToPascalCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToPascalCase with productionready", func(t *testing.T) {
		testString := "productionready"
		expected := "ProductionReady"
		resp := ToPascalCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToPascalCase with single word", func(t *testing.T) {
		testString := "hello"
		expected := "Hello"
		resp := ToPascalCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToPascalCase with empty string", func(t *testing.T) {
		testString := ""
		expected := ""
		resp := ToPascalCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToPascalCase with opensource", func(t *testing.T) {
		testString := "opensource"
		expected := "OpenSource"
		resp := ToPascalCase(testString)
		require.Equal(t, expected, resp)
	})
}
