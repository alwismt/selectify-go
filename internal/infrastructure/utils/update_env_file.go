package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UpdateEnvFile(key, value string) error {
	// Open .env file
	envFile, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer envFile.Close()

	// Scan .env file line by line
	scanner := bufio.NewScanner(envFile)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			line = key + "=" + value
		}
		lines = append(lines, line)
	}

	// Write updated lines to .env file
	if err := scanner.Err(); err != nil {
		return err
	}
	envFile.Truncate(0)
	envFile.Seek(0, 0)
	for _, line := range lines {
		if _, err := fmt.Fprintln(envFile, line); err != nil {
			return err
		}
	}

	return nil
}
