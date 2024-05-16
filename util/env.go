package util

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
			value = fallback
	}
	return value
}

func StringToIntWithDefault(str string, defaultValue int) int {
	value, err := strconv.Atoi(str)
	if err != nil {
			return defaultValue
	}
	return value
}