package utils

import (
	"bufio"
	"code2doc/internal/logger"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(path string) error {
	logger.Info(fmt.Sprintf("Loading environment variables from: %s", path))

	file, err := os.Open(path)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to open .env file: %v", err))
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	envCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		os.Setenv(key, val)
		envCount++
		logger.Debug(fmt.Sprintf("Loaded env: %s", key))
	}

	if err := scanner.Err(); err != nil {
		logger.Error(fmt.Sprintf("Error reading .env file: %v", err))
		return err
	}

	logger.Info(fmt.Sprintf("Successfully loaded %d environment variables", envCount))
	return nil
}
