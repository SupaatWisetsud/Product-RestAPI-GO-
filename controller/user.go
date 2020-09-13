package controller

import (
	"net/http"
	"product/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type UserController struct {
	DB *gorm.DB
}
type userResponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Avater string `json:"avater"`
	Role   string `json:"role"`
}

type userForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type userCreateResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"string"`
}

func (u *UserController) FindAll(ctx *gin.Context) {
	var user []models.User

	if err := u.DB.Find(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	serializedUser := []userReponse{}
	copier.Copy(&serializedUser, &user)
	ctx.JSON(http.StatusOK, gin.H{"users": serializedUser})
}

func (u *UserController) FindOne(ctx *gin.Context) {
	var user models.User

	if err := u.findOneByID(ctx, &user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	serializedUser := []userReponse{}
	copier.Copy(&serializedUser, &user)
	ctx.JSON(http.StatusOK, gin.H{"users": serializedUser})
}

func (u *UserController) Create(ctx *gin.Context) {
	var form userForm

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	copier.Copy(&user, &form)

	if err := u.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedUserCreateResponse userCreateResponse
	copier.Copy(&serializedUserCreateResponse, &user)
	ctx.JSON(http.StatusOK, gin.H{"user": serializedUserCreateResponse})
}

func (u *UserController) Delete(ctx *gin.Context) {
	var user models.User
	u.findOneByID(ctx, &user)

	if err := u.DB.Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (u *UserController) findOneByID(ctx *gin.Context, user *models.User) error {
	id := ctx.Param("id")

	if err := u.DB.First(user, id).Error; err != nil {
		return err
	}

	return nil
}
