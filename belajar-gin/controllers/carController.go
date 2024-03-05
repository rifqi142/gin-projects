package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID string `json:"car_id"`
	Brand string `json:"brand"`
	Price int    `json:"price"`
	Model string `json:"model"`
}

var CarDatas = []Car{}

// CreateCar is a function to create new car data
func CreateCar(ctx *gin.Context) {
	var newCar Car
	
	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newCar.CarID = fmt.Sprintf("c%d", len(CarDatas)+1)
	CarDatas = append(CarDatas, newCar)

	ctx.JSON(http.StatusCreated, gin.H {
		"car": newCar,
	})
}

// UpdateCar is a function to update car data
func UpdateCar (ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updatedCar Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			CarDatas[i] = updatedCar
			CarDatas[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Nout Found!!",
			"error_message": fmt.Sprintf("Car with ID %s not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("Car with ID %s has been updated", carID),
	})
}

// GetCar is a function to get car data
func GetCar (ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carData Car
	
	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carData = CarDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not Found!!",
			"error_message": fmt.Sprintf("Car with ID %s not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"car": carData,
	})
}

func DeleteCar (ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var carIndex int
	
	for i, car := range CarDatas {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not Found!!",
			"error_message": fmt.Sprintf("Car with ID %s not found", carID),
		})
		return
	}

	copy(CarDatas[carIndex:], CarDatas[carIndex+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("Car with ID %s has been deleted", carID),
	})
	
}