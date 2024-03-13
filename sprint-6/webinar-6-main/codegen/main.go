package main

import (
	"encoding/json"
	"fmt"
	"os"

	smartyapi "github.com/vmarunin/go-basic-2-webinar-6/codegen/model"
)

func main() {
	data, err := os.ReadFile("../examples/smarty.json")
	if err != nil {
		panic(err)
	}
	var smarty smartyapi.SmartyStreetsAPI
	err = json.Unmarshal(data, &smarty)
	if err != nil {
		panic(err)
	}

	fmt.Println(smarty[0].LastLine)
}
