package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Gamemaster struct {
	Pokemon []Pokemon `json:"pokemon"`
}

func main() {
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

	err = os.Mkdir("data", 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, p := range gm.Pokemon {
		b, err := json.Marshal(p.Spreads())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fpw, err := os.OpenFile(fmt.Sprintf("data/%s.json", p.ID), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = fpw.Write(b)
		fpw.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("files created")
}
