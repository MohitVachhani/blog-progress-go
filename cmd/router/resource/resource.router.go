package resourcerouter

import (
	resourceController "github.com/MohitVachhani/blog-progress/cmd/controller/resource"
	"github.com/gorilla/mux"
)

func InitResourceRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/get", resourceController.GetResource).Methods("GET")
	router.HandleFunc("/create", resourceController.CreateResource).Methods("POST")

	return router
}
