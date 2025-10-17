package strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToSnakeCase(t *testing.T) {
	t.Run("Test ToSnakeCase with helloworldfromme", func(t *testing.T) {
		testString := "helloworldfromme"
		expected := "hello_world_from_me"
		resp := ToSnakeCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToSnakeCase with productionready", func(t *testing.T) {
		testString := "productionready"
		expected := "production_ready"
		resp := ToSnakeCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToSnakeCase with single word", func(t *testing.T) {
		testString := "hello"
		expected := "hello"
		resp := ToSnakeCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToSnakeCase with empty string", func(t *testing.T) {
		testString := ""
		expected := ""
		resp := ToSnakeCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToSnakeCase with testcase", func(t *testing.T) {
		testString := "testcase"
		expected := "test_case"
		resp := ToSnakeCase(testString)
		require.Equal(t, expected, resp)
	})
	t.Run("Test ToCamelCase with opensource", func(t *testing.T) {
		testString := "opensource"
		expected := "open_source"
		resp := ToSnakeCase(testString)
		require.Equal(t, expected, resp)
	})
}
