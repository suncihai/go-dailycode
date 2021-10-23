package controllers

import (
	"encoding/json"
	"errors"
	"go-dailycode/db"
	"go-dailycode/models"
	"net/http"
)

// GetUsers godoc
// @Summary Get users request
// @Description Get users from dailycode database
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := []models.User{}
	
	rows, err := db.DB.Query("SELECT * FROM user")
	if err != nil {
		RespondWithError(err, w)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Password)
		if err != nil {
			RespondWithError(err, w)
		}
		users = append(users, user)
	}
    
	res := models.SuccessRes{Success: true, Status: 200, Data: users}
	RespondWithSuccess(res, w)
}

// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Register a new user"
// @Success 200
// @Router /register [post]
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		_, err = db.DB.Exec("INSERT INTO user (id, username, password) VALUES (? ? ?)", user.ID, user.Name, user.Password)
		if err != nil {
			RespondWithError(err, w)
		}
		RespondWithSuccess(true, w)
	}
}

// @Summary User login
// @Description User login function
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User login"
// @Success 200
// @Router /login [post]
func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    var password string 
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		RespondWithError(err, w)
	} else { 
		row := db.DB.QueryRow("SELECT * FROM user WHERE username = ?", user.Name)
		err = row.Scan(&password)
		if err != nil {
			RespondWithError(err, w)
		}
		if user.Password != password {
            err = errors.New("you input the wrong password")
			RespondWithError(err, w)
		}
		RespondWithSuccess(true, w)
	}
}