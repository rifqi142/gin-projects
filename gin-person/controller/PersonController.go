package controller

import (
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type personController struct {
}

func NewPersonController() *personController {
	return &personController{}
}

func (pc *personController) Create(ctx *gin.Context) {
	var newPerson model.Person

	err := ctx.ShouldBindJSON(&newPerson)
	if err != nil {
		var r model.Response = model.Response {
			Success: false,
			Error: err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	
	model.Persons = append(model.Persons, newPerson)

	var r model.Response = model.Response {
		Success: true,
		Data: newPerson,
	}
	ctx.JSON(http.StatusOK, r)
}

func (pc *personController) GetAll(ctx *gin.Context) {
	var r model.Response = model.Response {
		Success: true,
		Data: model.Persons,
	}
	ctx.JSON(http.StatusOK, r)
}