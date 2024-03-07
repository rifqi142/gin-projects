package controller

import (
	"fmt"
	"gin-person/model"
	"gin-person/repository"
	"gin-person/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type personController struct {
	personRepository repository.IPersonRepository
}

func NewPersonController(personRepository repository.IPersonRepository) *personController {
	return &personController{
		personRepository: personRepository,
	}
}

// Create new data person
func (pc *personController) Create(ctx *gin.Context) {
	var newPerson model.Person

	err := ctx.ShouldBindJSON(&newPerson)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	createdPerson, err := pc.personRepository.Create(newPerson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, createdPerson, ""))
}

// Get all data person
func (pc *personController) GetAll(ctx *gin.Context) {

	persons, err := pc.personRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, persons, ""))
}

// Update data person
func (pc *personController) Update(ctx *gin.Context) {
	var person model.Person

	err := ctx.ShouldBindJSON(&person)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	personID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}
	person.Id = personID

	updatedPerson, err := pc.personRepository.Update(person)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, updatedPerson, ""))
}

// Delete data person
func (pc *personController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	err = pc.personRepository.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, util.CreateResponse(true, fmt.Sprintf("Data dengan Id %d berhasil dihapus", id), ""))
}