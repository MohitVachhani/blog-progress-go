package Blockcontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	blockService "github.com/MohitVachhani/blog-progress/cmd/service/block"
	blockInterface "github.com/MohitVachhani/blog-progress/pkg/struct/block"
	"go.mongodb.org/mongo-driver/bson"
)

func GetBlock(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bson.M{"success": true, "Block": "Hello world"})
}

func CreateBlock(w http.ResponseWriter, r *http.Request) {

	// body parameters
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// transfer body to createBlock input
	var createBlockInput blockInterface.CreateBlockInput
	json.Unmarshal(body, &createBlockInput)

	blockService.CreateBlock(createBlockInput)

	// returns the client with json.
	json.NewEncoder(w).Encode(bson.M{"success": true})
}

func UpdateBlock(w http.ResponseWriter, r *http.Request) {

	// body parameters
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// transfer body to createBlock input
	var updateBlockInput blockInterface.UpdateBlockInput
	json.Unmarshal(body, &updateBlockInput)

	updatedBlock := blockService.UpdateBlock(updateBlockInput)

	// returns the client with json.
	json.NewEncoder(w).Encode(bson.M{"success": true, "block": updatedBlock})
}
