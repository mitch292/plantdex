package main

import (
	"log"
	"net/http"

	"github.com/mitch292/plantdex/plants"
)

func main() {
	plants := new(plants.Catalog)
	http.HandleFunc("/", plants.GetAllPlants)
	http.HandleFunc("/plant", plants.GetPlant)
	http.HandleFunc("/add", plants.AddPlant)
	log.Fatal(http.ListenAndServe("localhost:8088", nil))
}
