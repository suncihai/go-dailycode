package main

import (
	"encoding/json"
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
func getRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	records := []Record{}
	
	rows, err := db.Query("SELECT * FROM record")
	if err != nil {
		respondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var record Record
		err = rows.Scan(&record.ID, &record.Name, &record.Problem, &record.CreateAt, &record.Success, &record.Language)
		if err != nil {
			respondWithError(err, w)
		}
		records = append(records, record)
	}
    
	res := SuccessRes{Success: true, Status: 200, Data: records}
	respondWithSuccess(res, w)
}

// @Summary Create a new record
// @Description Create a new record with the input paylod
// @Tags records
// @Accept  json
// @Produce  json
// @Param record body Record true "Create record"
// @Success 200 {object} Record
// @Router /record [post]
func createRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	record := Record{}
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		respondWithError(err, w)
	} else { 
		_, err = db.Exec("INSERT INTO record (id, name, problem, create_at, success, language) VALUES (? ? ? ? ? ?)", record.ID, record.Name, record.Problem, record.CreateAt, record.Success, record.Language)
		if err != nil {
			respondWithError(err, w)
		}
		respondWithSuccess(true, w)
	}
}


