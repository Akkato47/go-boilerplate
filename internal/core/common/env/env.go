package env

import (
	"os"
	"strconv"
	"strings"
)

func GetString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func GetStrings(key, sep string, fallback []string) []string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return strings.Split(v, sep)
}

func GetInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	convertedV, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return convertedV
}
