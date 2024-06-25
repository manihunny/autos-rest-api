package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"net/http"
	_ "modernc.org/sqlite"
	"github.com/jmoiron/sqlx"
)

type Auto struct {
	ID    			string      `json:"id" db:"id"`
	Brand 			string    	`json:"brand" db:"brand"`
	Model  			string      `json:"model" db:"model"`
	Mileage    		float64     `json:"mileage" db:"mileage"`
	NumberOfOwners 	int    		`json:"number_of_owners" db:"number_of_owners"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /autos", getAllAutos)
	mux.HandleFunc("GET /autos/{id}", getAutoByID)
	mux.HandleFunc("POST /autos", createAuto)
	mux.HandleFunc("PUT /autos/{id}", updateAuto)
	mux.HandleFunc("PATCH /autos/{id}", partialUpdateAuto)
	mux.HandleFunc("DELETE /autos/{id}", deleteAuto)

	fmt.Println("Server started on port 80")
	http.ListenAndServe("localhost:80", mux)
}

func getAllAutos(w http.ResponseWriter, r *http.Request) {
	db, err := sqlx.Connect("sqlite", "store.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer db.Close()

	var autos []Auto
	err = db.Select(&autos, "SELECT * FROM autos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(autos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func getAutoByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	
	db, err := sqlx.Connect("sqlite", "store.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer db.Close()

	var auto Auto
	err = db.Get(&auto, "SELECT * FROM autos WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(auto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func createAuto(w http.ResponseWriter, r *http.Request) {
	var autoData Auto
	err := json.NewDecoder(r.Body).Decode(&autoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := sqlx.Connect("sqlite", "store.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer db.Close()

	_, err = db.NamedExec(`INSERT INTO autos (id, brand, model, mileage, number_of_owners) VALUES (:id, :brand, :model, :mileage, :number_of_owners)`, autoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(autoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func updateAuto(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	
	var autoData Auto
	err := json.NewDecoder(r.Body).Decode(&autoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db, err := sqlx.Connect("sqlite", "store.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE autos SET id = $1, brand = $2, model = $3, mileage = $4, number_of_owners = $5 WHERE id = $6`, autoData.ID, autoData.Brand, autoData.Model, autoData.Mileage, autoData.NumberOfOwners, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(autoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func partialUpdateAuto(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var autoData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&autoData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(autoData) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	query := "UPDATE autos SET "
	for key, value := range autoData {
		query += fmt.Sprintf(`%v = '%v',`, key, value)
	}
	query = strings.TrimRight(query, ",") + fmt.Sprintf(` WHERE id = '%v'`, id)

	db, err := sqlx.Connect("sqlite", "store.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteAuto(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db, err := sqlx.Connect("sqlite", "store.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer db.Close()

	_, err = db.Exec(`DELETE FROM autos WHERE id = $1`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}