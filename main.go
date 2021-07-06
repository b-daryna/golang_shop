package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Login string `json:"login"`
	Email string `json:"email"`
}

type Product struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Desc string `json:"desc"`
	Price float32 `json:"price"`
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to home page")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/products", returnAllProducts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnAllProducts(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Returning all Products:")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func main() {
//-------------------------CONNECTING TO OUR DB-----------------------------------

	Products = []Product{
		{Name: "Tomato", Type: "vegetable", Desc: "this is a red tomato", Price: 32.65},
		{Name: "Cucumber", Type: "vegetable", Desc: "this is a small cucmber", Price: 29.99},
		{Name: "Apple", Type: "fruit", Desc: "this is a green apple", Price: 15.40},
		{Name: "Blueberry", Type: "berry", Desc: "this is a blueberry", Price: 67},
	}
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

//------------------------DISPLAY ALL PRODUCTS IN CATEGORIES--------------------------------

	handleRequests()





//---------------TODO REGISTRATION-------------------------------------

//---------------CREATE TABLE(if not exists) AND INSERT USERS(if no duplicates)-----------------------------
	//table, err := db.Query("Create table users ")

	//insert, err := db.Query("insert into users (username) values ('JJJack')")
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()


	fmt.Printf("\nSuccessfully connected to database!\n")

}

const (
	host     = "localhost"
	port     = 5432
	user     = "dbessmer"
	password = "12345"
	dbname   = "shop"
)
