package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// Car represents a car entity.
type Car struct {
	ID           int
	Model        string
	Registration string
	Mileage      int
	Condition    bool // true for available, false for being rented
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "parkinglot.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable()

	r := mux.NewRouter()
	r.HandleFunc("/cars", ListAvailableCars).Methods("GET")
	r.HandleFunc("/cars", AddCar).Methods("POST")
	r.HandleFunc("/cars/{registration}/rentals", RentCar).Methods("POST")
	r.HandleFunc("/cars/{registration}/returns", ReturnCar).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func createTable() {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cars (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		model TEXT,
		registration TEXT,
		mileage INTEGER,
		condition BOOLEAN
	);
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func ListAvailableCars(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM cars WHERE condition = 1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var availableCars []Car

	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.Model, &car.Registration, &car.Mileage, &car.Condition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		availableCars = append(availableCars, car)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(availableCars)
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	var newCar Car
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existingCar := Car{}
	err := db.QueryRow("SELECT * FROM cars WHERE registration = ?", newCar.Registration).Scan(&existingCar.ID)
	if err == nil {
		http.Error(w, "Car with this registration already exists", http.StatusConflict)
		return
	}

	_, err = db.Exec("INSERT INTO cars (model, registration, mileage, condition) VALUES (?, ?, ?, ?)", newCar.Model, newCar.Registration, newCar.Mileage, newCar.Condition)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func RentCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	registration := vars["registration"]

	_, err := db.Exec("UPDATE cars SET condition = 0 WHERE registration = ?", registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ReturnCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	registration := vars["registration"]

	_, err := db.Exec("UPDATE cars SET condition = 1 WHERE registration = ?", registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
