package handler

import (
	"net/http"

	"github.com/Julioamoreno/proj-gopportunities/schemas"
	"github.com/Julioamoreno/proj-gopportunities/validator"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	updateOpeningRequest := validator.UpdateOpeningRequest{}

	ctx.BindJSON(&updateOpeningRequest)

	if err := updateOpeningRequest.Validate(); err != nil {
		logger.ErrorF("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, validator.IsRequiredError("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Opening not found")
		return
	}

	if updateOpeningRequest.Role != "" {
		opening.Role = updateOpeningRequest.Role
	}
	if updateOpeningRequest.Company != "" {
		opening.Company = updateOpeningRequest.Company
	}
	if updateOpeningRequest.Location != "" {
		opening.Location = updateOpeningRequest.Location
	}
	if updateOpeningRequest.Remote != nil {
		opening.Remote = *updateOpeningRequest.Remote
	}
	if updateOpeningRequest.Link != "" {
		opening.Link = updateOpeningRequest.Link
	}
	if updateOpeningRequest.Salary > 0 {
		opening.Salary = updateOpeningRequest.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error updating opening")
		return
	}

	sendSuccess(ctx, opening)
	return
}
