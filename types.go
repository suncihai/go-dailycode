package main

type Record struct {
	ID int `json:"id"`
	Name string `json:"username"`
	Problem int `json:"problem"`
	CreateAt string `json:"create_time"`
	Success bool `json:"sucess"`
	Language string `json:"lanauge"`
}

type User struct {
	ID int `json:"id"`
	Name string `json:"username"`
	Password string `json:"password"`
}