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
	var variable_name = os.Args[2]
	var variable_value = os.Args[3]

	// Load file into lines array.
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	lines := strings.Split(string(input), "\n")

	// Regular expression to find a valid configuration line.
	var regexstring = "^(\\s*)([a-zA-Z0-9_-]+)(\\s*)(=)(\\s*)(.*)"

	// Identify if some variable value was changed.
	var variable_changed = 0

	// Variable to set the identified separator ('=' or ' = ').
	var identified_separator = ""

	var found_variable_with_same_value = false

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
			if identified_separator == "" {
				identified_separator = matches[3] + "=" + matches[5]
			}

			// Change variable value.
			if strings.TrimSpace(matches[2]) == strings.TrimSpace(variable_name) && strings.TrimSpace(matches[6]) != strings.TrimSpace(variable_value) {
				/*
				   fmt.Printf("matches[0]: '%v'  type: %v\n", matches[0], reflect.TypeOf(matches[0]))
				   fmt.Printf("matches[1]: '%v'  type: %v\n", matches[1], reflect.TypeOf(matches[1]))
				   fmt.Printf("matches[2]: '%v'  type: %v\n", matches[2], reflect.TypeOf(matches[2]))
				   fmt.Printf("matches[3]: '%v'  type: %v\n", matches[3], reflect.TypeOf(matches[3]))
				   fmt.Printf("matches[4]: '%v'  type: %v\n", matches[4], reflect.TypeOf(matches[4]))
				   fmt.Printf("matches[5]: '%v'  type: %v\n", matches[5], reflect.TypeOf(matches[5]))
				*/

				linebufferwrite := matches[1] + variable_name + matches[3] + "=" + matches[5] + variable_value
				lines[i] = linebufferwrite
				variable_changed++
			}
			if strings.TrimSpace(matches[2]) == strings.TrimSpace(variable_name) && strings.TrimSpace(matches[6]) == strings.TrimSpace(variable_value) {
				found_variable_with_same_value = true
			}
		}
	}

	if variable_changed == 0 && !found_variable_with_same_value {
		// If no valid configuration line was found, set the separator ' = ' as default.
		if identified_separator == "" {
			identified_separator = " = "
		}

		// Add a new line to the end of file.
		lines = append(lines, variable_name+identified_separator+variable_value+"\n")
	}

	if !found_variable_with_same_value {
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
