package main

type Record struct {
	ID       int    `json:"id" example:"1"`
	Name     string `json:"username" example:"Peter"`
	Problem  int    `json:"problem" example:"1"`
	CreateAt string `json:"create_time" example:"2020-09-27"`
	Success  bool   `json:"sucess" example:"true"`
	Language string `json:"language" example:"English"`
}

type SuccessRes struct {
   Success bool `json:"success"`
   Status int64 `json:"status"`
   Data interface{} `json:"data"`
}

