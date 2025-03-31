package env

import (
	"errors"
	"os"
	"strconv"
)

// GetEnv is a helper function to retrieve the value of an environment variable.
// GetEnv retrieves the value of an environment variable.
// If the variable is not set, it returns the default value.
// The type of the default value determines the expected type of the environment variable.
// The function supports string, bool, int, and float64 types.
// If the conversion fails, it returns the default value and an error.
// Example usage: https://github.com/BenjaminNechicattu/gogenc/tree/main/examples
//
//	value, err := GetEnv("MY_ENV_VAR", "default")
//	if err != nil {
//
// handle error
//
//	}
//	fmt.Println(value) // prints the value of MY_ENV_VAR or "default" if not set
//
// python alternative is os.getenv("MY_ENV_VAR", "default"
func GetEnv[T string | bool | int | float64](key string, defaultValue T) (T, error) {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue, nil
	}

	switch any(defaultValue).(type) {
	case string:
		return any(value).(T), nil
	case bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue, err
		}
		return any(b).(T), nil
	case int:
		i, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue, err
		}
		return any(i).(T), nil
	case float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return defaultValue, err
		}
		return any(f).(T), nil
	default:
		return defaultValue, errors.New("unsupported type")
	}
}
