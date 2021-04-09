package main

import (
	"fmt"
	"io/ioutil"
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
	myRouter.HandleFunc("/api/v1/arithmetic", countThings).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func returnStatus200(w http.ResponseWriter, r *http.Request) {

}

func countThings(w http.ResponseWriter, r *http.Request) {

	log.Info("Request count things")
	log.Info("Request = ", r.Body)
	log.Info("Method = ", r.Method)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	bodyStr := string(body)
	if len(bodyStr) == 0 {
		fmt.Fprintf(w, "4")
	}
	expression, err := govaluate.NewEvaluableExpression(bodyStr)
	if err != nil {
		log.Info(err)
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Info(err)
	}
	log.Info(result)
	log.Info("result = ", result)
	fmt.Fprintf(w, "4")
}

func main() {
	log.SetOutput(os.Stdout)
	log.Info("Rest API v2.0 - Mux Routers")

	handleRequests()
}
