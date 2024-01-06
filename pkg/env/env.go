// File: env.go
// Functionalities:
//	- Retrieve environment variables from config files.

package env

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var defaultValues = map[string][]string{
	"SERVER_HOST": {"localhost"},
	"SERVER_PORT": {"7171"},
}

func GetConfig(name string) []string {
	defaultConfigDir := os.Getenv("CONFIG_DIR")
	if defaultConfigDir == "" {
		defaultConfigDir = "configs/default.yaml"
	}

	configs, err := godotenv.Read(defaultConfigDir)
	if err != nil {
		return defaultValues[name]
	}

	if configs["CUSTOM_CONFIG_FILE"] != "" {
		// fmt.Println("Custom config file specified. Using custom config file", configs["CUSTOM_CONFIG_FILE"])
		configs, err = godotenv.Read("configs/" + configs["CUSTOM_CONFIG_FILE"])

		if err != nil {
			// fmt.Println("Error reading custom config file: " + err.Error() + "\nUsing default config file.")
			configs, err = godotenv.Read(defaultConfigDir)

			if err != nil {
				// fmt.Println("Error reading default config file: " + err.Error() + "\nUsing default values.")
				return defaultValues[name]
			}
		}
	}

	return strings.FieldsFunc(
		configs[name],
		func(r rune) bool {
			return r == ','
		})
}
