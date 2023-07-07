package handler

import (
	"fmt"
	"net/http"

	"github.com/Julioamoreno/proj-gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Error listing openings"))
		return
	}

	sendSuccess(ctx, openings)
}
