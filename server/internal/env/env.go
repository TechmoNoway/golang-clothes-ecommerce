package env

import (
	"fmt"
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {

	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	fmt.Println(val)

	return val
}

func GetInt(key string, fallback int) int {

	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
