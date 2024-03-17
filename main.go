package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Bar struct {
	Id string `json:"id"`
}
type Any struct {
	Data []Bar `json:"data"`
}

func main() {
	files, err := ioutil.ReadDir("in/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		fileContent, err := os.ReadFile(fmt.Sprintf("in/%s", file.Name()))
		if err != nil {

		}
		var data Any
		err = json.Unmarshal(fileContent, &data)
		fmt.Println(data)
		fmt.Print(data.Data[0].Id)
		if err != nil {

		}
	}

}
