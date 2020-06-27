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
	Name          string `json:"name" bson:"name"`
	Size          string `json:"size" bson:"size"`
	WaterSchedule int    `json:"waterSchedule" bson:"waterSchedule"` // TODO: consts / map to seconds
	SunLevel      int    `json:"sunLevel" bson:"sunLevel"`           // TODO: consts for this
	Notes         string `json:"notes" bson:"notes"`
	IsPetSafe     bool   `json:"isPetSafe" bson:"isPetSafe"`
	Food          int    `json:"food" bson:"food"` // TODO: consts / map to seconds
	ShouldMist    bool   `json:"shouldMist" bson:"shouldMist"`
}

func (p Plant) String() string {
	return p.Name
}

// GetPlant will return a plant from mongo db
func (c Catalog) GetPlant(w http.ResponseWriter, req *http.Request) {
	// plantName := req.URL.Query().Get("name")

	plant, err := getPlantByName("Test")

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
		Name:          "Test",
		Size:          "M",
		WaterSchedule: 13400,
		SunLevel:      3,
		Notes:         "Nothing special to do here.",
		IsPetSafe:     true,
		Food:          260000,
		ShouldMist:    true,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Failure adding the plant %s\n", err)
		return
	}

	fmt.Fprintf(w, "We added a plant!")
}
