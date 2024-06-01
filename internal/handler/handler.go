package handler

import (
	"database/sql"
	"emission-tracker-server/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRecords handles the GET /records route.
func GetRecords(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, datetime, user_id, category, subcategory, amount, unit_type, input_datetime FROM record")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var records []model.Record
		for rows.Next() {
			var r model.Record
			if err := rows.Scan(&r.ID, &r.Name, &r.Datetime, &r.UserID, &r.Category, &r.Subcategory, &r.Amount, &r.UnitType, &r.InputDatetime); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			records = append(records, r)
		}

		c.JSON(http.StatusOK, records)
	}
}

// GetRecordByID handles the GET /records/:id route.
func GetRecordByID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var r model.Record
		err := db.QueryRow("SELECT id, name, datetime, user_id, category, subcategory, amount, unit_type, input_datetime FROM record WHERE id = ?", id).Scan(
			&r.ID, &r.Name, &r.Datetime, &r.UserID, &r.Category, &r.Subcategory, &r.Amount, &r.UnitType, &r.InputDatetime)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "record not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		c.JSON(http.StatusOK, r)
	}
}

// PostRecord handles the POST /records route.
func PostRecord(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newRecord model.Record

		if err := c.BindJSON(&newRecord); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := db.Exec("INSERT INTO record (name, datetime, user_id, category, subcategory, amount, unit_type, input_datetime) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			newRecord.Name, newRecord.Datetime, newRecord.UserID, newRecord.Category, newRecord.Subcategory, newRecord.Amount, newRecord.UnitType, newRecord.InputDatetime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newRecord.ID = int(id)
		c.JSON(http.StatusCreated, newRecord)
	}
}
