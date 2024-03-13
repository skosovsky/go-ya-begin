package main

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	data, err := os.ReadFile("../examples/smarty.json")
	if err != nil {
		panic(err)
	}
	var smarty SmartyStreetsAPI
	err = json.Unmarshal(data, &smarty)
	if err != nil {
		panic(err)
	}

	data, err = yaml.Marshal(&smarty)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))
	// fmt.Println(string(smarty[0].Components))
}
