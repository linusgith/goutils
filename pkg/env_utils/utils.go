package goutils

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

type logStruct struct{}

func Log() LogStruct {
	return logStruct{}
}

type noLogStruct struct{}

func NoLog() NoLogStruct {
	return noLogStruct{}
}

// ParseEnvDurationDefault returns the value of the environment variable 'env' parsed as a time.Duration.
// If the variable is not set or cannot be parsed, it returns the provided durationDefault.
// Logs a warning if parsing fails.
func (l *Log) ParseEnvDurationDefault(env string, durationDefault time.Duration, logger *zap.Logger) time.Duration {

	timeout := os.Getenv(env)

	duration, err := time.ParseDuration(timeout)
	if err != nil {
		logger.Warn("Could not parse from .env, setting default", zap.String("variable", env), zap.Duration("default", durationDefault), zap.Error(err))
		return durationDefault
	}

	return duration
}

// ParseEnvDurationPanic returns the value of the environment variable 'env' parsed as a time.Duration.
// If the variable is not set or cannot be parsed, it logs a panic with the provided logger.
func (l *Log) ParseEnvDurationPanic(env string, logger *zap.Logger) time.Duration {

	timeout := os.Getenv(env)

	duration, err := time.ParseDuration(timeout)
	if err != nil {
		logger.Panic("Could not parse from .env", zap.String("variable", env), zap.Error(err))
	}

	return duration
}

// ParseEnvIntDefault returns the value of the environment variable 'env' parsed as an int.
// If the variable is not set or cannot be parsed, it returns the provided intDefault.
// Logs a warning if parsing fails and logs the variable name and value at debug level.
func (l *Log) ParseEnvIntDefault(env string, intDefault int, logger *zap.Logger) int {

	intString := os.Getenv(env)

	logger.Debug("got env variable", zap.String("variable", env), zap.String("value", intString))

	parsedInt, err := strconv.Atoi(intString)
	if err != nil {
		logger.Warn("Could not parse from .env, setting default", zap.String("variable", env), zap.Int("default", intDefault), zap.Error(err))
		return intDefault
	}

	return parsedInt
}

// ParseEnvIntPanic returns the value of the environment variable 'env' parsed as an int.
// If the variable is not set or cannot be parsed, it logs a panic with the provided logger.
func (l *Log) ParseEnvIntPanic(env string, logger *zap.Logger) int {

	intString := os.Getenv(env)

	parsedInt, err := strconv.Atoi(intString)
	if err != nil {
		logger.Panic("Could not parse from .env", zap.String("variable", env), zap.Error(err))
	}

	return parsedInt
}

// ParseEnvStringDefault returns the value of the environment variable 'env'.
// If the variable is not set or is empty, it returns the provided defaultValue.
// Logs the variable name and value at debug level.
func (l *Log) ParseEnvStringDefault(env string, defaultValue string, logger *zap.Logger) string {
	envString := os.Getenv(env)

	logger.Debug("got env variable", zap.String("variable", env), zap.String("value", envString))

	if envString == "" {
		return defaultValue
	}

	return envString

}

// ParseEnvStringPanic returns the value of the environment variable 'env'.
// If the variable is not set or is empty, it logs a panic with the provided logger.
// Logs the variable name and value at debug level.
func (l *Log) ParseEnvStringPanic(env string, logger *zap.Logger) string {
	envString := os.Getenv(env)

	logger.Debug("got env variable", zap.String("variable", env), zap.String("value", envString))

	if envString == "" {
		logger.Panic("environment variable was empty")
	}

	return envString

}

func (NoLog) ParseEnvDurationDefault(env string, durationDefault time.Duration) time.Duration {

	timeout := os.Getenv(env)

	duration, err := time.ParseDuration(timeout)
	if err != nil {
		return durationDefault
	}

	return duration
}

// ParseEnvDurationPanic returns the value of the environment variable 'env' parsed as a time.Duration.
// If the variable is not set or cannot be parsed, it logs a panic with the provided logger.
func (NoLog) ParseEnvDurationPanic(env string) time.Duration {

	timeout := os.Getenv(env)

	duration, err := time.ParseDuration(timeout)
	if err != nil {
		panic(fmt.Sprintf("could not parse from .env: %s, %v", env, err))
	}

	return duration
}

// ParseEnvIntDefault returns the value of the environment variable 'env' parsed as an int.
// If the variable is not set or cannot be parsed, it returns the provided intDefault.
// Logs a warning if parsing fails and logs the variable name and value at debug level.
func (NoLog) ParseEnvIntDefault(env string, intDefault int) int {

	intString := os.Getenv(env)

	parsedInt, err := strconv.Atoi(intString)
	if err != nil {
		return intDefault
	}

	return parsedInt
}

// ParseEnvIntPanic returns the value of the environment variable 'env' parsed as an int.
// If the variable is not set or cannot be parsed, it logs a panic with the provided logger.
func (NoLog) ParseEnvIntPanic(env string) int {

	intString := os.Getenv(env)

	parsedInt, err := strconv.Atoi(intString)
	if err != nil {
		panic(fmt.Sprintf("could not parse from .env: %s, %v", env, err))
	}

	return parsedInt
}

// ParseEnvStringDefault returns the value of the environment variable 'env'.
// If the variable is not set or is empty, it returns the provided defaultValue.
// Logs the variable name and value at debug level.
func (NoLog) ParseEnvStringDefault(env string, defaultValue string) string {
	envString := os.Getenv(env)

	if envString == "" {
		return defaultValue
	}

	return envString

}

// ParseEnvStringPanic returns the value of the environment variable 'env'.
// If the variable is not set or is empty, it logs a panic with the provided logger.
// Logs the variable name and value at debug level.
func (NoLog) ParseEnvStringPanic(env string) string {
	envString := os.Getenv(env)

	if envString == "" {
		panic(fmt.Sprintf("could not parse from .env: %s, %w", env, errors.New("env was empty")))
	}

	return envString

}
