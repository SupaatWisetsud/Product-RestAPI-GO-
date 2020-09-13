package controller

import (
	"mime/multipart"
	"net/http"
	"os"
	"product/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

type productForm struct {
	Name       string                `form:"name" binding:"required"`
	Price      uint                  `form:"price"`
	CategoryID uint                  `form:"category_id" binding:"required"`
	Image      *multipart.FileHeader `form:"image"`
}

type productResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Image string `json:"image"`
}

func (p *ProductController) FindAll(ctx *gin.Context) {
	var product []models.Product

	if err := p.DB.Find(&product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	serialzedProduct := []productResponse{}
	copier.Copy(&serialzedProduct, &product)

	ctx.JSON(http.StatusOK, gin.H{"products": serialzedProduct})
}

func (p *ProductController) FindOne(ctx *gin.Context) {
	var product models.Product

	p.findByID(ctx, &product)

	var serializedProduct productResponse
	copier.Copy(&serializedProduct, &product)

	ctx.JSON(http.StatusOK, gin.H{"product": serializedProduct})
}

func (p *ProductController) Create(ctx *gin.Context) {

	var form productForm

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	copier.Copy(&product, &form)

	if err := p.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	p.setImageProduct(ctx, &product)

	var serializedProduct productResponse
	copier.Copy(&serializedProduct, &product)
	ctx.JSON(http.StatusOK, gin.H{"product": serializedProduct})
}

func (p *ProductController) Update(ctx *gin.Context) {

}

func (p *ProductController) Delete(ctx *gin.Context) {
	var product models.Product

	p.findByID(ctx, &product)

	if err := p.DB.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (p *ProductController) setImageProduct(ctx *gin.Context, product *models.Product) {

	file, _ := ctx.FormFile("image")

	if file == nil {
		return
	}

	if product.Image != "" {
		pwd, _ := os.Getwd()
		product.Image = strings.Replace(product.Image, os.Getenv("HOST"), "", 1)

		os.Remove(pwd + string(product.Image))
	}

	dir := "uploads/products/" + strconv.Itoa(int(product.ID))

	os.MkdirAll(dir, 0755)

	dir = dir + "/" + file.Filename

	if err := ctx.SaveUploadedFile(file, dir); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	product.Image = os.Getenv("HOST") + "/" + dir

	if err := p.DB.Save(product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
}

func (p *ProductController) findByID(ctx *gin.Context, product *models.Product) {

	id := ctx.Param("id")

	if err := p.DB.First(product, id).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
}
