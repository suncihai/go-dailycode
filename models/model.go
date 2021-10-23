package models

import "database/sql"

type Record struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"username" example:"Peter"`
	Problem  int    `json:"problem" example:"1"`
	CreateAt string `json:"create_time" example:"2020-09-27"`
	Success  bool   `json:"sucess" example:"true"`
	Language string `json:"language" example:"English"`
}
type User struct {
	ID int `json:"id" example:"1"`
	Name string `json:"usernmae" example:"Peter"`
	Password string `json:"password" example:"xxxxxxxxx"`
}

type Problem struct {
	ID int `json:"id" example:"1"`
	Number int `json:"number" example:"1"`
	Name string `json:"name" exampe:"Two Sum"`
	Difficulty string `json:"difficulty" example:"Easy"`
	Tag string `json:"tag" example:"Array Linkedlist"`
	Category sql.NullString `json:"category" example:"Google Facebook"`
}
type SuccessRes struct {
	Success bool `json:"success"`
	Status int64 `json:"status"`
	Data interface{} `json:"data"`
 }