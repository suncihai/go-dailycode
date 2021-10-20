package main

import (
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

