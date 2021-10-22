package controllers

import (
	"encoding/json"
	"go-dailycode/db"
	"go-dailycode/models"
	"net/http"
)

// GetRecords godoc
// @Summary Get records request
// @Description Get records from dailycode database
// @Tags records
// @Accept  json
// @Produce  json
// @Success 200 {array} Record
// @Router /records [get]
func GetRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	records := []models.Record{}
	
	rows, err := db.DB.Query("SELECT * FROM record")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var record models.Record
		err = rows.Scan(&record.ID, &record.Name, &record.Problem, &record.CreateAt, &record.Success, &record.Language)
		if err != nil {
			RespondWithError(err, w)
		}
		records = append(records, record)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: records}
	RespondWithSuccess(res, w)
}

// @Summary Create a new record
// @Description Create a new record with the input paylod
// @Tags records
// @Accept  json
// @Produce  json
// @Param record body Record true "Create record"
// @Success 200 {object} Record
// @Router /record [post]
func CreateRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	record := models.Record{}
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		_, err = db.DB.Exec("INSERT INTO record (id, name, problem, create_at, success, language) VALUES (? ? ? ? ? ?)", record.ID, record.Name, record.Problem, record.CreateAt, record.Success, record.Language)
		if err != nil {
			RespondWithError(err, w)
		}
		RespondWithSuccess(true, w)
	}
}