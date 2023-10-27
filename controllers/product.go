package controllers

import (
	"errors"
	"my-gin-gorm/config"
	"my-gin-gorm/models"
	"my-gin-gorm/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(ctx *gin.Context) {
	var products []models.Product
	err := config.DB.Debug().Find(&products).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": err})
		return
	}
	ctx.JSON(http.StatusOK, &products)
}

func CreateProduct(ctx *gin.Context) {
	var product models.Product
	// validation
	if err := ctx.ShouldBind(&product); err != nil {
		validation := utils.ParseError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": validation,
		})
		return

		// for _, fieldErr := range err.(validator.ValidationErrors) {
		// 	field := strings.ToLower(string(fieldErr.StructField()))
		// 	// tag := fieldErr.ActualTag()
		// 	// errStr := fmt.Sprintf("The %s error in %s", field, tag)
		// 	errSplit := strings.Split(fieldErr.Error(), ":")
		// 	errStr := strings.Replace(errSplit[2], fieldErr.StructField(), field, -1)

		// 	ctx.JSON(http.StatusBadRequest, gin.H{
		// 		"error": errStr,
		// 	})
		// 	return // exit on first error
		// }
	}

	if err := config.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, &product)
}

func GetProductByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var product models.Product
	if result := config.DB.Where("id = ?", id).First(&product); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func UpdateProduct(ctx *gin.Context) {
	var product models.Product
	id := ctx.Params.ByName("id")

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		errorCode := http.StatusNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = http.StatusNotFound
		}
		ctx.AbortWithStatusJSON(errorCode, gin.H{"error": err.Error()})
		return
	}

	// Validate input
	var input models.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&product).Updates(input).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(ctx *gin.Context) {
	var product models.Product
	id := ctx.Params.ByName("id")
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		var errorCode int = http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorCode = http.StatusNotFound
		}
		ctx.AbortWithStatusJSON(errorCode, gin.H{"error": err.Error()})
		return
	}

	// soft delete
	// config.DB.Delete(&models.Product{}, id)
	// config.DB.Where("id = ?", id).Delete(&models.Product{})
	// config.DB.Delete(&product)

	config.DB.Unscoped().Delete(&product) // permanet delete
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
