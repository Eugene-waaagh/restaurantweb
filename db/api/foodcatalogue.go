package api

import (
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

	foodcatalogue, err := server.store.CreateCategory(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, foodcatalogue)
}
