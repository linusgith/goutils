package goutils

import (
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

// ParseEnvDurationDefault returns the value of the environment variable 'env' parsed as a time.Duration.
// If the variable is not set or cannot be parsed, it returns the provided durationDefault.
// Logs a warning if parsing fails.
func ParseEnvDurationDefault(env string, durationDefault time.Duration, logger *zap.Logger) time.Duration {

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
func ParseEnvDurationPanic(env string, logger *zap.Logger) time.Duration {

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
func ParseEnvIntDefault(env string, intDefault int, logger *zap.Logger) int {

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
func ParseEnvIntPanic(env string, logger *zap.Logger) int {

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
func ParseEnvStringDefault(env string, defaultValue string, logger *zap.Logger) string {
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
func ParseEnvStringPanic(env string, logger *zap.Logger) string {
	envString := os.Getenv(env)

	logger.Debug("got env variable", zap.String("variable", env), zap.String("value", envString))

	if envString == "" {
		logger.Panic("environment variable was empty")
	}

	return envString

}
