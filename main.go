package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Archetype struct {
	Name     string
	Matchers []string
}
type Pokemon struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Attacks []Attack `json:"attacks"`
}

type Attack struct {
	Text string `json:"text"`
}

type PokemonIn struct {
	Data []Pokemon `json:"data"`
}

var archeTypes = []Archetype{
	{
		Name: "Confuse Attack",
		Matchers: []string{
			"the Defending Pokémon is now Confused.",
		},
	},
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
		var data PokemonIn
		err = json.Unmarshal(fileContent, &data)
		transform(data.Data)
		if err != nil {

		}
	}

}

func transform(data []Pokemon) {
	for _, pokemon := range data {
		analyseArchetype((pokemon))
	}
	toWrite, _ := json.Marshal(data)
	_ = ioutil.WriteFile("output.json", toWrite, 0644)
}

func analyseArchetype(pokemon Pokemon) {
	for _, attack := range pokemon.Attacks {
		for _, archetype := range archeTypes {
			for _, matcher := range archetype.Matchers {
				if strings.Contains(attack.Text, matcher) {
					fmt.Println(attack)
				}
			}
		}
	}
}
