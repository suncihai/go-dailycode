package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/go-openapi/runtime/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
    db, err = sql.Open("mysql", ConnectionString)
	if(err != nil) {
		panic(err.Error())
	}
	fmt.Println("successfully connected to database.")

	router := mux.NewRouter()
	setupRoutesForRecords(router)
    
    router.Handle("/swagger.yaml", http.FileServer(http.Dir("/")))

	opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil);
	router.Handle("/docs", sh)
    
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr: port,
		WriteTimeout:  15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
	//http.ListenAndServe(port, router)
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Home Page!")
}