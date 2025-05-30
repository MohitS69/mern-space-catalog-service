package main

import (
	"catalog/config"
	"catalog/handler"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	conf, err := config.LoadConfig("./config/development.yaml")
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "api is healthy")
	})
	client, err := mongo.Connect(options.Client().ApplyURI(conf.DB.URI))
	if err != nil {
		conf.Logger.Panic(err)
	}
	handler.SetupeCategoryRoutes(router.Group("/category"), client)
	router.Run(fmt.Sprintf(":%d", conf.Server.Port))
}
