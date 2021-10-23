package routers

import (
	"go-dailycode/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutesForRecords(router *mux.Router) {
	enableCORS(router)
    
	// router.HandleFunc("/", main.home)
	router.HandleFunc("/records", controllers.GetRecords).Methods("GET")
	router.HandleFunc("/record", controllers.CreateRecord).Methods("POST")
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/problems", controllers.GetProblems).Methods("GET")
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Contro-Allow-Origin", AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}

