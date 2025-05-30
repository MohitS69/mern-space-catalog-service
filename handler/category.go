package handler

import (
	"catalog/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoryHandler struct {
	db *mongo.Database
}

func SetupeCategoryRoutes(router *gin.RouterGroup, db *mongo.Database) {
	handler := CategoryHandler{
		db,
	}
	router.POST("/", handler.create)
}

func (h *CategoryHandler) create(c *gin.Context) {
	var payload models.Category
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coll := models.GetCategoryCollection(h.db)
	res, err := coll.InsertOne(c.Request.Context(), payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
