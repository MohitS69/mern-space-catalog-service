package main

import (
	"catalog/config"
	"catalog/handler"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig("./config/development.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize MongoDB client
	client, err := mongo.Connect(options.Client().ApplyURI(conf.DB.URI))
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Ensure connection is valid by pinging primary
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("MongoDB disconnect failed: %v", err)
		}
	}()

	// Setup router
	router := gin.Default()

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "API is healthy")
	})

	// Setup routes for category
	handler.SetupeCategoryRoutes(router.Group("/category"), client.Database("catalog"))

	// Run server
	addr := fmt.Sprintf(":%d", conf.Server.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

