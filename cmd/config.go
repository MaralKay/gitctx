package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func saveConfig(name, gitConfigPath string, sshConfigPath string) error {
	file, err := os.OpenFile(configFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%s=%s=%s\n", name, gitConfigPath, sshConfigPath)
	if err != nil {
		return err
	}

	return nil
}

func readConfig() (map[string][]string, error) {
	config := make(map[string][]string)

	file, err := os.Open(configFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return config, nil // File does not exist, return empty config
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := parseConfigLine(line)
		if len(parts) == 3 {
			var pathsArray []string
			pathsArray = append(pathsArray, parts[1], parts[2])
			config[parts[0]] = pathsArray
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func parseConfigLine(line string) []string {
	parts := strings.SplitN(line, "=", 3)
	if len(parts) == 3 {
		return parts
	}
	return []string{}
}
