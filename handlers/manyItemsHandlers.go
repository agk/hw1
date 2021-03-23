package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"hw1/models"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllItems(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get infos about all items in database")
	if len(models.DB)>0 {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(models.DB)
	} else {
		writer.WriteHeader(403)
		msg := models.Message{ Message: "Error: No one items found in store back" }
		json.NewEncoder(writer).Encode(msg)
	}
}
