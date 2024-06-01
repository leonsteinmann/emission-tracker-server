package scripts

import (
	"database/sql"
	"emission-tracker-server/pkg/model"
	"fmt"
	"math/rand"
	"time"
)

func InsertNewEntries(db *sql.DB, numEntries int) error {
	fmt.Println("Trying to add 100 entries")
	// Prepare statement for inserting data
	stmt, err := db.Prepare("INSERT INTO record (name, datetime, user_id, category_id, subcategory_id, amount, unit_type, input_datetime) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Generate X new random entries
	for i := 0; i < numEntries; i++ {
		// Generate random data for the entry
		name := fmt.Sprintf("Entry%d", i)
		datetime := time.Now().Add(-time.Duration(rand.Intn(30*24*60*60)) * time.Second) // Random datetime within the last 30 days

		// Randomly select a category
		category := model.Categories[rand.Intn(len(model.Categories))]

		// Generate a random subcategory (placeholder for actual logic)
		subcategory := 0

		// Generate a random amount
		amount := rand.Float64() * 100

		// Insert the entry into the database
		_, err := stmt.Exec(name, datetime, nil, category.ID, subcategory, amount, "CO2e", time.Now())
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
