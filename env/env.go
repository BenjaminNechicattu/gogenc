package env

import (
	"errors"
	"os"
	"strconv"
)

// GetEnv retrieves the value of an environment variable.
// If the variable is not set, it returns the default value.
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
