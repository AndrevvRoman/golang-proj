package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Knetic/govaluate"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Info("Endpoint Hit: homePage")
}

func handleRequests() {
	port, exists := os.LookupEnv("PORT")
	if exists == false {
		fmt.Printf("NO PORT")
	}
	log.Info("Founded port ", port)
	//http.ListenAndServe(":"+port, nil)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/v1/health", returnStatus200).Methods(http.MethodHead)
	myRouter.HandleFunc("api/v1/arithmetic", countThings).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func returnStatus200(w http.ResponseWriter, r *http.Request) {

}

func countThings(w http.ResponseWriter, r *http.Request) {

	log.Info("Request count things")
	log.Info("Request = ", r.Body)
	log.Info("Method = ", r.Method)
	expression, err := govaluate.NewEvaluableExpression("2 + 3 - 1")
	if err != nil {
		log.Info(err)
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Info(err)
	}
	log.Info(result)
	// json.NewEncoder(w).Encode(expression)
}

func main() {
	log.SetOutput(os.Stdout)
	log.Info("Rest API v2.0 - Mux Routers")

	handleRequests()
}
