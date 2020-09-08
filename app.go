package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize method to create connect db
func (a *App) Initialize(user, password, dbname string) {

	connectString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	fmt.Println(connectString)
	var err error
	a.DB, err = sql.Open("postgres", connectString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run method to start server
func (a *App) Run(addr string) {}
