package env_test

import (
	"os"
	"testing"

	"github.com/BenjaminNechicattu/gogenc/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEnv(t *testing.T) {
	t.Run("TestGetEnvString", func(t *testing.T) {
		os.Setenv("TEST_ENV", "test_value")
		val, err := env.GetEnv("TEST_ENV", "default")
		require.NoError(t, err)
		assert.Equal(t, "test_value", val)
	})
	t.Run("TestGetEnvStringDefault", func(t *testing.T) {
		val, err := env.GetEnv("NON_EXISTENT_ENV", "default")
		require.NoError(t, err)
		assert.Equal(t, "default", val)
	})
	t.Run("TestGetEnvInt", func(t *testing.T) {
		os.Setenv("TEST_ENV_INT", "42")
		val, err := env.GetEnv("TEST_ENV_INT", 0)
		require.NoError(t, err)
		assert.Equal(t, 42, val)
	})
	t.Run("TestGetEnvIntDefault", func(t *testing.T) {
		val, err := env.GetEnv("NON_EXISTENT_ENV_INT", 42)
		require.NoError(t, err)
		assert.Equal(t, 42, val)
	})
	t.Run("TestGetEnvBool", func(t *testing.T) {
		os.Setenv("TEST_ENV_BOOL", "true")
		val, err := env.GetEnv("TEST_ENV_BOOL", false)
		require.NoError(t, err)
		assert.Equal(t, true, val)
	})
	t.Run("TestGetEnvBoolDefault", func(t *testing.T) {
		val, err := env.GetEnv("NON_EXISTENT_ENV_BOOL", false)
		require.NoError(t, err)
		assert.Equal(t, false, val)
	})
}
