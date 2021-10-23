package controllers

import (
	"go-dailycode/db"
	"go-dailycode/models"
	"net/http"
)

// GetProblems godoc
// @Summary Get problems request
// @Description Get problems from dailycode database
// @Tags problems
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Problem
// @Router /problems [get]
func GetProblems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	problems := []models.Problem{}
	
	rows, err := db.DB.Query("SELECT * FROM problem")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var problem models.Problem
		err = rows.Scan(&problem.ID, &problem.Number, &problem.Name, &problem.Difficulty, &problem.Tag, &problem.Category)
		if err != nil {
			RespondWithError(err, w)
		}
		problems = append(problems, problem)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: problems}
	RespondWithSuccess(res, w)
}