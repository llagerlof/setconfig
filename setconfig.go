package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	// File name must be specified.
	if len(os.Args) < 2 {
		fmt.Println("Missing filename.")
		os.Exit(1)
	}

	// The file name must be passed as argument.
	if len(os.Args) < 3 {
		fmt.Println("Missing variable name.")
		os.Exit(1)
	}

	// The variable value must be passed as argument.
	if len(os.Args) < 4 {
		fmt.Println("Missing variable value.")
		os.Exit(1)
	}

	// File must exist.
	if !fileExists(os.Args[1]) {
		fmt.Println("File not exist.")
		os.Exit(1)
	}

	// Store arguments in variables.
	var filename = os.Args[1]
	var variableName = os.Args[2]
	var variableValue = os.Args[3]

	// Load file into lines array.
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")

	// Regular expression to find a valid configuration line.
	var regexstring = "^(\\s*)([a-zA-Z0-9_-]+)(\\s*)(=)(\\s*)(.*)"

	// Identify if some variable value was changed.
	var variableChanged = 0

	// Variable to set the identified separator ('=' or ' = ').
	var identifiedSeparator = ""

	var foundVariableWithSameValue = false

	// Loop in array, changing variables as needed.
	for i, line := range lines {
		fmt.Println(line)

		regex := regexp.MustCompile(regexstring)
		matches := regex.FindStringSubmatch(line)
		// Group 0: full match
		// Group 1: space(s) or nothing
		// Group 2: variable name
		// Group 3: space(s) or nothing
		// Group 4: equal
		// Group 5: space(s) or nothing
		// Group 6: variable value

		// A valid 'variable = value' line
		if len(matches) >= 5 {
			// Identify default separator.
			if identifiedSeparator == "" {
				identifiedSeparator = matches[3] + "=" + matches[5]
			}

			// Change variable value.
			if strings.TrimSpace(matches[2]) == strings.TrimSpace(variableName) && strings.TrimSpace(matches[6]) != strings.TrimSpace(variableValue) {
				linebufferwrite := matches[1] + variableName + matches[3] + "=" + matches[5] + variableValue
				lines[i] = linebufferwrite
				variableChanged++
			}
			if strings.TrimSpace(matches[2]) == strings.TrimSpace(variableName) && strings.TrimSpace(matches[6]) == strings.TrimSpace(variableValue) {
				foundVariableWithSameValue = true
			}
		}
	}

	if variableChanged == 0 && !foundVariableWithSameValue {
		// If no valid configuration line was found, set the separator ' = ' as default.
		if identifiedSeparator == "" {
			identifiedSeparator = " = "
		}

		// Add a new line to the end of file.
		lines = append(lines, variableName+identifiedSeparator+variableValue+"\n")
	}

	if !foundVariableWithSameValue {
		// Change the variable value
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(filename, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		} else {
			fmt.Printf("File saved.\n")
		}
	} else {
		fmt.Printf("No changes were made.\n")
	}

	os.Exit(0)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
