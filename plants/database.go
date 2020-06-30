package plants

import (
	"database/sql"
	"log"

	// This driver needs to be made available to the database/sql package
	_ "github.com/mattn/go-sqlite3"
)

// DB handles our interaction with the sqlite db
var db *sql.DB

// InitDB will create our SQLite3 DB if it doesn't exist
func InitDB() {
	database, err := sql.Open("sqlite3", "./plantdex.db")
	if err != nil {
		log.Fatalf("Error opening the sqlite connection: %s\n", err)
	}

	stmt, err := database.Prepare("CREATE TABLE IF NOT EXISTS plants (id INTEGER PRIMARY KEY, name TEXT, size INTEGER, water_schedule INTEGER, sun_level INTEGER, notes TEXT, is_pet_safe INTEGER, food INTEGER, should_mist INTEGER)")
	if err != nil {
		log.Fatalf("Error generating the plants table create if not exist statment %s\n", err)
	}
	stmt.Exec()

	db = database
}

// addPlantToDB will add a plant to our plants table in the database
func addPlantToDB(p *Plant) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO plants (name, size, water_schedule, sun_level, notes, is_pet_safe, food, should_mist) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Error preparing the insert statement: %s\n", err)
		return 0, err
	}

	res, err := stmt.Exec(p.Name, p.Size, p.WaterSchedule, p.SunLevel, p.Notes, p.IsPetSafe, p.Food, p.ShouldMist)
	if err != nil {
		log.Fatalf("Error adding the plant to the database: %s\n", err)
		return 0, err
	}

	newPlantID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("Error readingt he last inserted item's ID: %s\n", err)
	}

	return newPlantID, nil
}

func getPlantFromDB(id int64) (*Plant, error) {
	var p Plant
	err := db.QueryRow("SELECT * FROM plants WHERE id = ?", id).Scan(&p.Id, &p.Name, &p.Size, &p.WaterSchedule, &p.SunLevel, &p.Notes, &p.IsPetSafe, &p.Food, &p.ShouldMist)
	if err != nil {
		log.Fatalf("Error fetching the plant from the database: %s\n", err)
		return &p, err
	}

	return &p, nil
}

func getAllPlantsFromDB() (*Plants, error) {
	var plants Plants
	stmt, err := db.Prepare("SELECT * FROM plants")
	if err != nil {
		log.Fatalf("Error preparing the all plants query: %s\n", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatalf("Error excecuting the all plants query: %s\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		var p Plant
		err := rows.Scan(&p.Id, &p.Name, &p.Size, &p.WaterSchedule, &p.SunLevel, &p.Notes, &p.IsPetSafe, &p.Food, &p.ShouldMist)
		if err != nil {
			log.Fatalf("Error while scanning the rows: %s\n", err)
		}
		plants.Catalog = append(plants.Catalog, &p)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("Error in our final check: %s\n", err)
	}

	return &plants, nil

}
