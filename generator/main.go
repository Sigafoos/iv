package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Sigafoos/iv/model"
)

const dataPath = "../data"

type Gamemaster struct {
	Pokemon []Pokemon `json:"pokemon"`
}

func main() {
	log.Print("opening gamemaster...")
	fp, err := os.Open("gamemaster.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(fp)
	fp.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var gm Gamemaster
	err = json.Unmarshal(body, &gm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Print("removing existing data...")
	// don't care about an error as we're removing it if it exists
	_ = os.RemoveAll(dataPath)
	err = os.MkdirAll(dataPath+"/great", 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = os.Mkdir(dataPath+"/ultra", 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, p := range gm.Pokemon {
		log.Printf("calculating spread for %s...", p.Name)
		err := savePokemon(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	log.Println("all files created")
}

func savePokemon(p Pokemon) error {
	great, ultra := p.Spreads()
	err := saveSpread(great, fmt.Sprintf("%s/great/%s.json", dataPath, p.ID))
	if err != nil {
		return err
	}
	err = saveSpread(ultra, fmt.Sprintf("%s/ultra/%s.json", dataPath, p.ID))
	return err
}

func saveSpread(spread map[string]model.Spread, path string) error {
	b, err := json.Marshal(spread)
	if err != nil {
		return err
	}
	fpw, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = fpw.Write(b)
	fpw.Close()
	return err
}
