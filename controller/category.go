package controller

import (
	"net/http"
	"product/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

type categoryRespones struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Product []struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Price uint   `json:"price"`
	} `json:"product"`
}

type categoryForm struct {
	Name string `json:"name" bindding:"required"`
}

type categoryUpdate struct {
	Name string `json:"name"`
}

func (c *CategoryController) FindAll(ctx *gin.Context) {
	var category []models.Category

	if err := c.DB.Preload("Product").Find(&category).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory []categoryRespones

	copier.Copy(&serializedCategory, &category)

	ctx.JSON(http.StatusOK, gin.H{"categorys": serializedCategory})

}
func (c *CategoryController) FindOne(ctx *gin.Context) {
	var category models.Category

	err := c.findOneByID(&category, ctx)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory categoryRespones

	copier.Copy(&serializedCategory, &category)

	ctx.JSON(http.StatusOK, gin.H{"category": serializedCategory})
}

func (c *CategoryController) Create(ctx *gin.Context) {
	var form categoryForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	copier.Copy(&category, &form)

	if err := c.DB.Create(&category).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory categoryRespones
	copier.Copy(&serializedCategory, &category)
	ctx.JSON(http.StatusOK, gin.H{"category": serializedCategory})
}

func (c *CategoryController) Update(ctx *gin.Context) {
	var form categoryUpdate

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var category models.Category

	if err := c.findOneByID(&category, ctx); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Model(&category).Update(&form).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedCategory categoryRespones

	copier.Copy(&serializedCategory, &category)

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"category": serializedCategory})
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	var category models.Category

	err := c.findOneByID(&category, ctx)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if err := c.DB.Unscoped().Delete(&category).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *CategoryController) findOneByID(category *models.Category, ctx *gin.Context) error {
	id := ctx.Param("id")

	if err := c.DB.Preload("Product").First(category, id).Error; err != nil {
		return err
	}

	return nil
}
