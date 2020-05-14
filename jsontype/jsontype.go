package jsontype

import (
	"fmt"
	"reflect"
)

func typeof(result interface{}) string {
	return reflect.TypeOf(result).String()
}

func itArray(json []interface{}, scope int) int {
	for _, v := range json {
		FindType(v, scope)
		fmt.Println("")
	}

	return scope
}

func itMap(json map[string]interface{}, scope int) int {
	for k, v := range json {
		insertTab(scope)
		fmt.Print(k, ": ")
		FindType(v, scope)
	}

	return scope
}

func insertTab(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print("   ") // 3 spaces
	}
}

// FindType : find the type of a json object
func FindType(json interface{}, level int) {
	switch typeof(json) {
	case "[]interface {}":
		level = itArray(json.([]interface{}), level+1)
	case "map[string]interface {}":
		fmt.Print("\n")
		level = itMap(json.(map[string]interface{}), level+1)
	case "interface {}":
		fmt.Println("is object")
	case "string":
		fmt.Println(json.(string))
	case "float64":
		fmt.Println(json.(float64))
	case "int":
		fmt.Println(json.(int))
	default:
		fmt.Println("could not find type ", typeof(json))
	}
}
