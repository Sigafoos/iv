package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/Sigafoos/iv/model"
)

var (
	re       = regexp.MustCompile(`^\d{1,2}/\d{1,2}/\d{1,2}$`)
	stripper = regexp.MustCompile(`[\.()]`)
)

const (
	LeagueGreat  = "great"
	LeagueUltra  = "ultra"
	LeagueMaster = "master"
)

var cache map[string]map[string]map[string]model.Spread

func main() {
	cache = make(map[string]map[string]map[string]model.Spread)
	cache[LeagueGreat] = make(map[string]map[string]model.Spread)
	cache[LeagueUltra] = make(map[string]map[string]model.Spread)
	cache[LeagueMaster] = make(map[string]map[string]model.Spread)

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
	league := r.FormValue("league")
	if league == "" {
		league = LeagueGreat
	}

	if league != LeagueGreat && league != LeagueUltra && league != LeagueMaster {
		// TODO more than this
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	name := filename(pokemon)

	s, ok := cache[league][name]
	if !ok {
		var spread map[string]model.Spread
		fp, err := os.Open(fmt.Sprintf("data/%s/%s.json", league, name))
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
		cache[league][name] = spread
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

func filename(pokemon string) string {
	pokemon = strings.ToLower(pokemon)
	pokemon = strings.Replace(pokemon, " ", "_", -1)
	pokemon = stripper.ReplaceAllString(pokemon, "")
	return pokemon
}
