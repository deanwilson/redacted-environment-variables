package main

import (
	"fmt"
	"os"
	"strings"
)

func ToRedact() []string {
	var redactedNames = []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"GITHUB_TOKEN",
		"GITHUB_AUTH_TOKEN",
		"_TOKEN_",
	}

	return redactedNames
}

func main() {
	redactedVariables := ToRedact()

	for _, envvar := range os.Environ() {
		pair := strings.SplitN(envvar, "=", 2)

		name, value := pair[0], pair[1]

		for _, redacted := range redactedVariables {
			if strings.Contains(name, redacted) {
				value = "XXXXXXXX"
			}
		}

		fmt.Println(name + "=" + value)
	}
}
