package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"hw1/models"
	"github.com/gorilla/mux"
)

func GetItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	Item, ok := models.FindItemById(id)
	log.Println("Get Item with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Error: Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(Item)
	}
}

func CreateItem(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new Item ....")
	var Item models.Item
	var ok bool

	err := json.NewDecoder(request.Body).Decode(&Item)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	var newItemID int = len(models.DB) + 1
	Item.ID = newItemID
	models.DB = append(models.DB, Item)
	
	_, ok = models.FindItemById(Item.ID)
	if (!ok) {
		writer.WriteHeader(201)
		msg := models.Message{Message: "Item created"}
		json.NewEncoder(writer).Encode(msg)
	} else {
		writer.WriteHeader(400)
		msg := models.Message{Message: "Error: Ityem with that id already exists"}
		json.NewEncoder(writer).Encode(msg)
	}
}

func UpdateItemById(writer http.ResponseWriter, request *http.Request) {
	//var ok bool
	initHeaders(writer)
	log.Println("Updating Item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Error: do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindItemById(id)
	var newItem models.Item
	if !ok {
		log.Println("Item not found in data base . id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "Error: Item with that id not found"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&newItem)
	if err != nil {
		msg := models.Message{Message: "Error: provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	//TODO:Нужно заменить oldItem на newItem в DB!

	models.DB = models.FindAndReplaceItemById(id, newItem)
	writer.WriteHeader(202)
	json.NewEncoder(writer).Encode(newItem)
	return
	
}

func DeleteItemById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Deleting Item ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("error while parsing happend:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Error: do not use parameter ID as uncasted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	_, ok := models.FindItemById(id)
	if !ok {
		log.Println("Item not found in database. id :", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "Error: Item with that ID does not exists in database"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	//TODO: Нужно удалить Item из DB
	
	msg := models.Message{Message: "successfully deleted requested item"}
	json.NewEncoder(writer).Encode(msg)
}
