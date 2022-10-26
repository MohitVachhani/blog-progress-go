package main

import (
	"context"
	"log"
	"net/http"
	"time"

	blockRouter "github.com/MohitVachhani/blog-progress/cmd/router/block"
	resourceRouter "github.com/MohitVachhani/blog-progress/cmd/router/resource"
	envUtil "github.com/MohitVachhani/blog-progress/pkg/utils/env"
	mongoUtils "github.com/MohitVachhani/blog-progress/pkg/utils/mongo"

	"github.com/gorilla/mux"
)

func initializeRoutes() {

	// init router
	var router = mux.NewRouter()

	// initialize all the entities router

	// Resource api endpoints
	resourceR := router.PathPrefix("/resource").Subrouter()
	resourceRouter.InitResourceRouter(resourceR)

	// block api endpoints
	blockR := router.PathPrefix("/block").Subrouter()
	blockRouter.InitBlockRouter(blockR)

	// start server and throw error if anything goes wrong.
	port := ":" + envUtil.Get("PORT")
	log.Fatal(http.ListenAndServe(port, router))

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	mongoClient := mongoUtils.MongoConnection(ctx)
	defer mongoClient.Disconnect(ctx)

	// initialize routes
	initializeRoutes()
}
