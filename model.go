package main

type Record struct {
	ID int `json:"id"`
	Name string `json:"username"`
	Problem int `json:"problem"`
	CreateAt string `json:"create_time"`
	Success bool `json:"sucess"`
	Language string `json:"language"`
}

type SuccessRes struct {
   Success bool `json:"success"`
   Status int64 `json:"status"`
   Data interface{} `json:"data"`
}

