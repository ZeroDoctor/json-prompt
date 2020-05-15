package jsontype

import (
	"fmt"
	"reflect"
)

func typeof(result interface{}) string {
	return reflect.TypeOf(result).String()
}

func itArray(json []interface{}, scope int) int {
	insertTab(scope - 1)
	fmt.Println("[")
	for _, v := range json {
		insertTab(scope)
		FindType(v, scope)
	}
	insertTab(scope - 1)
	fmt.Println("]")
	return scope
}

func itMap(json map[string]interface{}, sortedKeys []string, scope int) int {
	insertTab(scope - 1)
	fmt.Println("{")
	for k := range sortedKeys {
		insertTab(scope)
		fmt.Print("\"", sortedKeys[k], "\": ")
		FindType(json[sortedKeys[k]], scope)
	}
	insertTab(scope - 1)
	fmt.Println("}")
	return scope
}

func insertTab(amount int) {
	for i := 0; i < amount; i++ {
		fmt.Print("  ") // 2 spaces
	}
}

// FindType : find the type of a json object
func FindType(json interface{}, level int) {
	switch typeof(json) {
	case "[]interface {}":
		fmt.Println("")
		level = itArray(json.([]interface{}), level+1)
	case "map[string]interface {}":
		fmt.Println("")
		jsonMap := json.(map[string]interface{})
		level = itMap(jsonMap, SortMap(jsonMap), level+1)
	case "string":
		fmt.Printf("\"%s\",\n", json.(string))
	case "float64":
		fmt.Printf("%g,\n", json.(float64))
	case "int":
		fmt.Printf("%d,\n", json.(int))
	case "interface {}":
		fmt.Println("is object")
	default:
		fmt.Println("could not find type ", typeof(json))
	}
}
