package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("../examples/compose.json")
	if err != nil {
		panic(err)
	}
	var compose map[string]interface{}
	err = json.Unmarshal(data, &compose)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", compose)

	services := compose["services"].(map[string]interface{})
	proxy := services["proxy"].(map[string]interface{})
	dependsOn := proxy["depends_on"].([]interface{})
	if dependsOn[0].(string) == "backend" {
		fmt.Println("Hurray!")
	}

	// switch proxy["depends_on"].(type) {
	// switch compose["services"].(type) {
	switch dependsOn[0].(type) {
	case []interface{}:
		fmt.Println("Array")
	case map[string]interface{}:
		fmt.Println("Map")
	case string:
		fmt.Println("String")
	}
}
