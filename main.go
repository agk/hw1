package main

import (
	"log"
	"net/http"
	"os"

	"hw1/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                    string
	itemResourcePrefix      string = apiPrefix + "/item"  //api/v1/item/
	manyItemsResourcePrefix string = apiPrefix + "/items" //api/v1/items/
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildItemResource(router, itemResourcePrefix)
	utils.BuildManyItemsResourcePrefix(router, manyItemsResourcePrefix)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
