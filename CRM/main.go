package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

func setHeader(response http.ResponseWriter) {
	response.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
}

func setSuccessSatus(response http.ResponseWriter) {
	response.WriteHeader(http.StatusOK)
}

// index
func index(response http.ResponseWriter, request *http.Request) {
	setHeader(response)
	setSuccessSatus(response)
	json.NewEncoder(response).Encode(CRM_MESSAGE)
}

// Get all customers
func getCustomers(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(mock_database)
}

// Get a Customer by id
func getCustomer(response http.ResponseWriter, request *http.Request) {
	setHeader(response)
	id := mux.Vars(request)["id"]
	if _, ok := mock_database[id]; ok {
		setSuccessSatus(response)
		json.NewEncoder(response).Encode(mock_database[id])
	} else {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(FAIL_MESSAGE)
	}
}

// Handle to create Customer
func addCustomer(response http.ResponseWriter, request *http.Request) {
	setHeader(response)
	var customer customer
	data, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(data, &customer)

	if _, ok := mock_database[customer.ID]; ok {
		response.WriteHeader(http.StatusConflict)
		json.NewEncoder(response).Encode(FAIL_MESSAGE)
	} else {
		mock_database[customer.ID] = customer
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(SUCCESS_MESSAGE)
	}
}

// Handle to update Customer
func updateCustomer(response http.ResponseWriter, request *http.Request) {
	var updateCustomer customer
	data, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(data, &updateCustomer)

	setHeader(response)
	id := mux.Vars(request)["id"]
	if _, ok := mock_database[id]; ok {
		mock_database[id] = updateCustomer
		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(SUCCESS_MESSAGE)
	} else {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(FAIL_MESSAGE)
	}

}

// Handle to delete Customer
func deleteCustomer(response http.ResponseWriter, request *http.Request) {
	setHeader(response)
	id := mux.Vars(request)["id"]

	if _, ok := mock_database[id]; ok {
		delete(mock_database, id)
		setSuccessSatus(response)
		json.NewEncoder(response).Encode(SUCCESS_MESSAGE)
	} else {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode(FAIL_MESSAGE)
	}
}

func main() {
	fmt.Println("Server is starting at " + PORT)
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	fmt.Println("Server is running at " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
