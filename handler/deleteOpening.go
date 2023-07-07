package handler

import (
	"fmt"
	"net/http"

	"github.com/Julioamoreno/proj-gopportunities/schemas"
	"github.com/Julioamoreno/proj-gopportunities/validator"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, validator.IsRequiredError("id",
			"queryParameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening id: %v, not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		logger.ErrorF("Error deleting opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening id: %v", id))
		return
	}

	sendSuccess(ctx, opening)
}
