package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/Sigafoos/iv/model"
)

var (
	re = regexp.MustCompile(`^\d{1,2}/\d{1,2}/\d{1,2}$`)
)
var list map[string]map[string]model.Spread

func main() {
	list = make(map[string]map[string]model.Spread)
	http.HandleFunc("/iv", serveIV)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("server running on port " + port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}

func serveIV(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("accept") != "application/json" {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	pokemon := r.FormValue("pokemon")
	ivs := r.FormValue("ivs")
	if pokemon == "" || ivs == "" || !re.MatchString(ivs) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s, ok := list[pokemon]
	if !ok {
		var spread map[string]model.Spread
		fp, err := os.Open(fmt.Sprintf("data/%s.json", pokemon))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		body, err := ioutil.ReadAll(fp)
		fp.Close()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(body, &spread)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		list[pokemon] = spread
		s = spread
	}
	response, ok := s[ivs]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(b))
}
