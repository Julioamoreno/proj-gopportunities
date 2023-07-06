package handler

import (
	"net/http"

	"github.com/Julioamoreno/proj-gopportunities/schemas"
	"github.com/Julioamoreno/proj-gopportunities/validator"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	createOpeningRequest := validator.CreateOpeningRequest{}

	ctx.BindJSON(&createOpeningRequest)

	if err := createOpeningRequest.Validate(); err != nil {
		logger.ErrorF("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     createOpeningRequest.Role,
		Company:  createOpeningRequest.Company,
		Location: createOpeningRequest.Location,
		Remote:   *createOpeningRequest.Remote,
		Link:     createOpeningRequest.Link,
		Salary:   createOpeningRequest.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorF("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "Error creating opening on database")
		return
	}

	sendSuccess(ctx, opening)
	return
}
