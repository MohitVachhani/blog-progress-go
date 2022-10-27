package blockrouter

import (
	blockController "github.com/MohitVachhani/blog-progress/cmd/controller/block"
	"github.com/gorilla/mux"
)

func InitBlockRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/create", blockController.CreateBlock).Methods("POST")
	router.HandleFunc("/update", blockController.UpdateBlock).Methods("POST")
	return router
}
