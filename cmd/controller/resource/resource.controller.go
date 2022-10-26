package resourcecontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	resourceService "github.com/MohitVachhani/blog-progress/cmd/service/resource"
	resourceInterface "github.com/MohitVachhani/blog-progress/pkg/struct/resource"
	"go.mongodb.org/mongo-driver/bson"
)

func GetResource(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(bson.M{"success": true, "resource": "Hello world"})
}

func CreateResource(w http.ResponseWriter, r *http.Request) {

	// body parameters
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// transfer body to createResource input
	var createResourceInput resourceInterface.CreateResourceInput
	json.Unmarshal(body, &createResourceInput)

	createdResource := resourceService.CreateResource(createResourceInput)

	// returns the client with json.
	json.NewEncoder(w).Encode(bson.M{"success": true, "resource": createdResource})
}
