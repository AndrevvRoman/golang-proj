package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	port, exists := os.LookupEnv("PORT")
	if exists == false {
		fmt.Printf("NO PORT")
	}

	//http.ListenAndServe(":"+port, nil)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/v1/health", returnStatus200)
	myRouter.HandleFunc("/api/v1/arithmetic", countThings)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func returnStatus200(w http.ResponseWriter, r *http.Request) {

}

func countThings(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request count things")
	fmt.Println("Request = ", r.Body)
	// expression, err := govaluate.NewEvaluableExpression("2 + 3 - 1")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// result, err := expression.Evaluate(nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(result)
	// json.NewEncoder(w).Encode(expression)
}

func main() {

	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
