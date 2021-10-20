build:
	swag init -g main.go
	go build main.go

run:
	swag init -g main.go
	go run main.go

