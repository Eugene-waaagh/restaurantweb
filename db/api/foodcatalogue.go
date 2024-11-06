package api

import (
	db "github.com/eugene-waaagh/restaurantweb/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createCatalogueRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createCatalogue(ctx *gin.Context) {
	var req createCatalogueRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFoodParams{
		Name: req.Name,
	}

	foodc
}
