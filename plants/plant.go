package plants

import (
	"fmt"
	"net/http"
)

// Catalog - the holder of all plants
type Catalog struct {
	plants []*Plant
}

// Plant is a description of a single house plant
type Plant struct {
	name          string
	size          string
	waterSchedule int // TODO: consts / map to seconds
	sunLevel      int // TODO: consts for this
	notes         string
	isPetSafe     bool
	food          int // TODO: consts / map to seconds
	shouldMist    bool
}

func (p Plant) String() string {
	return p.name
}

// GetPlant will return a plant from mongo db
func (c Catalog) GetPlant(w http.ResponseWriter, req *http.Request) {
	plantName := req.URL.Query().Get("name")

	plant, err := getPlantByName(plantName)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Did not find what you're looking for")
		return
	}

	fmt.Fprintf(w, "%s\n", plant)
}

// GetAllPlants will return all the plants we have in our db
func (c Catalog) GetAllPlants(w http.ResponseWriter, req *http.Request) {
	plants, err := getAllPlants()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Did not find what you're looking for")
		return
	}

	for _, plant := range plants {
		fmt.Fprintf(w, "%s\n", plant)
	}

}

// AddPlant will add a plant to our db store
func (c Catalog) AddPlant(w http.ResponseWriter, req *http.Request) {
	_, err := insertPlant(Plant{
		name:          "Test",
		size:          "M",
		waterSchedule: 13400,
		sunLevel:      3,
		notes:         "Nothing special to do here.",
		isPetSafe:     true,
		food:          260000,
		shouldMist:    true,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failure adding the plant %s\n", err)
		return
	}

	fmt.Fprintf(w, "We added a plant!")
}
