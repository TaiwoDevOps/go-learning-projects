package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseConfig(content string) (map[string]string, error) {
	config := make(map[string]string)

	re := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)

	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		matches := re.FindStringSubmatch(trimmed)

		if matches == nil {
			fmt.Printf("invalid config line %d\n", lineNo)
			continue
		}

		name, value1, value2, value3 := matches[1], matches[2], matches[3], matches[4]
		if value1 != "" {
			config[name] = value1
		} else if value2 != "" {
			config[name] = value2
		} else if value3 != "" {
			config[name] = value3
		}

	}

	return config, nil
}

func main() {
	envConfigs := `
		# Application Xonfiguration
App_NAME="My Cool App"
APP_VERSION="1.0.2-beta" # Version with quotes
PORT=8080
DEBUG_MODE="true"
# Database settings
DB_HOST=localhost
DB_USER = admin
DB_PASSWORD = "p@s$w ord with Sp@ces!" # Quoted password
API_ENDPOINT = https://api.example.com

# An empty value 
EMPTY_KEY=
ANOTHER_KEY_NO_VALUE =
	`

	config, err := parseConfig(envConfigs)

	if err != nil {
		fmt.Printf("Code exited with error")
		os.Exit(1)
	}

	for key, value := range config {
		fmt.Printf("%s=%q\n", key, value)
	}
}

// key := matches[1]
// var value string

// if matches[2] != "" {
// 	value = matches[2]
// } else if matches[3] != "" {
// 	value = matches[3]
// } else {
// 	value = matches[4]
// }

// config[key] = value
