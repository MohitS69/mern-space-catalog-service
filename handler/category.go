package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type CategoryHandler struct {
    db *mongo.Client
}

func SetupeCategoryRoutes(router *gin.RouterGroup,db *mongo.Client){
    handler := CategoryHandler{
        db,
    }
    router.POST("/",handler.create)
}

func (h *CategoryHandler) create(c *gin.Context){
    c.String(http.StatusOK,"hello")
}
