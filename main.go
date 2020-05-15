package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zerodoctor/json-prompt/jsontype"
)

func getInput() ([]byte, error) {
	args := os.Args[1:]
	if len(args) <= 0 {
		info, err := os.Stdin.Stat()

		if err != nil {
			panic(err)
		}
		if info.Mode()&os.ModeNamedPipe == 0 {
			fmt.Println("Usage: <json-output> | json-prompt [<json-hardcoded>]") // change this later
			return []byte(""), errors.New("Error: unexpected format")
		}

		output, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Failed to read input")
		}

		return output, nil
	}

	return []byte(args[0]), nil
}

func main() {

	fmt.Println("formatting... ")

	jsonInput, err := getInput()
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}

	var resultMap interface{}
	err = json.Unmarshal(jsonInput, &resultMap)

	if err != nil {
		fmt.Println("ERROR:", err)
	}

	jsontype.FindType(resultMap, 0)
}
