package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type Superhero struct {
	Name       string     `json:"name"`
	Biography  Biography  `json:"biography"`
	Powerstats Powerstats `json:"powerstats"`
	Images     Images     `json:"images"`
}

type Biography struct {
	FullName string `json:"fullName"`
}

type Powerstats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

type Images struct {
	XS string `json:"xs"`
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
}

var superheroes = map[string]Superhero{
	"wolverine": {
		Name: "Wolverine",
		Biography: Biography{
			FullName: "John Logan",
		},
		Powerstats: Powerstats{
			Intelligence: 63,
			Strength:     32,
			Speed:        50,
			Durability:   100,
			Power:        89,
			Combat:       100,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
		},
	},
	"spiderman": {
		Name: "Spider-Man",
		Biography: Biography{
			FullName: "Peter Parker",
		},
		Powerstats: Powerstats{
			Intelligence: 90,
			Strength:     55,
			Speed:        67,
			Durability:   75,
			Power:        74,
			Combat:       85,
		},
		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spiderman.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spiderman.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spiderman.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spiderman.jpg",
		},
	},
	"ironman": {
		Name: "Iron Man",
		Biography: Biography{
			FullName: "Tony Stark",
		},
		Powerstats: Powerstats{
			Intelligence: 100,
			Strength:     85,
			Speed:        58,
			Durability:   85,
			Power:        100,
			Combat:       64,
		},

		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
		},
	},
	"blackwidow": {
		Name: "Black Widow",
		Biography: Biography{
			FullName: "Natasha Romanoff",
		},
		Powerstats: Powerstats{
			Intelligence: 75,
			Strength:     13,
			Speed:        33,
			Durability:   30,
			Power:        36,
			Combat:       100,
		},

		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/107-black-widow.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/107-black-widow.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/107-black-widow.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/107-black-widow.jpg",
		},
	},
	"thor": {
		Name: "Thor",
		Biography: Biography{
			FullName: "Thor Odinson",
		},
		Powerstats: Powerstats{
			Intelligence: 69,
			Strength:     100,
			Speed:        83,
			Durability:   100,
			Power:        100,
			Combat:       100,
		},

		Images: Images{
			XS: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/659-thor.jpg",
			SM: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/659-thor.jpg",
			MD: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/659-thor.jpg",
			LG: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/659-thor.jpg",
		},
	},
}

func getSuperhero(w http.ResponseWriter, r *http.Request) {
	heroName := r.URL.Query().Get("hero")
	heroName = strings.ToLower(heroName)
	heroName = strings.ReplaceAll(heroName, " ", "")
	if hero, ok := superheroes[heroName]; ok {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hero)
	} else {
		http.Error(w, "Superhero not found", http.StatusNotFound)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/superhero", getSuperhero).Methods("GET")

	http.ListenAndServe(":8080", r)
}
