package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zerodoctor/json-prompt/jsontype"
)

func getInput() (string, error) {
	args := os.Args[1:]
	if len(args) <= 0 {
		info, err := os.Stdin.Stat()

		if err != nil {
			panic(err)
		}
		if info.Mode()&os.ModeNamedPipe == 0 {
			fmt.Println("Usage: <json-output> | json-prompt [<json-hardcoded>]") // change this later
			return "", errors.New("Error: unexpected format")
		}

		output, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Failed to read input")
		}

		return string(output), nil
	}

	return string(args[0]), nil
}

func main() {

	fmt.Println("formatting... ")

	jsonInput, err := getInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	var resultMap interface{}
	err = json.Unmarshal([]byte(jsonInput), &resultMap)

	if err != nil {
		fmt.Println(err)
	}

	jsontype.FindType(resultMap, 0)
}
