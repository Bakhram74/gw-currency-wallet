package env

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetDomain -Получение домена из адреса пользователя
func GetDomain(v string) string {
	r := strings.Split(v, "@")
	return r[1]
}

// GetEnv - Получение переменной окружения с типом STRING
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvAsTime(name string, defaultValue time.Duration) time.Duration {
	valueStr := GetEnv(name, "")

	if valueStr == "" {
		return defaultValue
	}
	if value, err := time.ParseDuration(valueStr); err == nil {
		return value
	} else {
		log.Fatalf("GetEnvAsBool error: %v", err)
	}
	return defaultValue
}

// GetEnvAsInt - Получение переменной окружения с типом INT
func GetEnvAsInt(name string, defaultValue int) int {
	valueStr := GetEnv(name, "")

	if valueStr == "" {
		return defaultValue
	}

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	} else {
		log.Fatalf("GetEnvAsInt error: %v", err)
	}

	return defaultValue
}

// GetEnvAsBool - Получение переменной окружения с типом BOOL
func GetEnvAsBool(name string, defaultValue bool) bool {
	valueStr := GetEnv(name, "")

	if valueStr == "" {
		return defaultValue
	}

	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	} else {
		log.Fatalf("GetEnvAsBool error: %v", err)
	}

	return defaultValue
}

func GetEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := GetEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
