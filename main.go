package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	_ "go-dailycode/docs"

	_ "github.com/go-sql-driver/mysql"

	httpSwagger "github.com/swaggo/http-swagger"

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
    
	// opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	// sh := middleware.SwaggerUI(opts, nil);
	// router.Handle("/docs", sh)
	// router.Handle("/swagger.yaml", http.FileServer(http.Dir("/")))
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
    
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

var orders []Order
var prevOrderID = 0

type Order struct {
	OrderID      string    `json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"Leo Messi"`
	OrderedAt    time.Time `json:"orderedAt" example:"2019-11-09T21:21:46+00:00"`
	Items        []Item    `json:"items"`
}

// Item represents the model for an item in the order
type Item struct {
	ItemID      string `json:"itemId" example:"A1B2C3"`
	Description string `json:"description" example:"A random description"`
	Quantity    int    `json:"quantity" example:"1"`
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the input paylod
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body Order true "Create order"
// @Success 200 {object} Order
// @Router /orders [post]
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	json.NewDecoder(r.Body).Decode(&order)
	prevOrderID++
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)
}

// GetOrders godoc
// @Summary Get details of all orders
// @Description Get details of all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} Order
// @Router /orders [get]
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

// GetOrder godoc
// @Summary Get details for a given orderId
// @Description Get details of order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order"
// @Success 200 {object} Order
// @Router /orders/{orderId} [get]
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for _, order := range orders {
		if order.OrderID == inputOrderID {
			json.NewEncoder(w).Encode(order)
			return
		}
	}
}

// UpdateOrder godoc
// @Summary Update order identified by the given orderId
// @Description Update the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be updated"
// @Success 200 {object} Order
// @Router /orders/{orderId} [post]
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			var updatedOrder Order
			json.NewDecoder(r.Body).Decode(&updatedOrder)
			orders = append(orders, updatedOrder)
			json.NewEncoder(w).Encode(updatedOrder)
			return
		}
	}
}

// DeleteOrder godoc
// @Summary Delete order identified by the given orderId
// @Description Delete the order corresponding to the input orderId
// @Tags orders
// @Accept  json
// @Produce  json
// @Param orderId path int true "ID of the order to be deleted"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputOrderID := params["orderId"]
	for i, order := range orders {
		if order.OrderID == inputOrderID {
			orders = append(orders[:i], orders[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}