package controller

import (
	"net/http"
	"product/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

type authForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type authReponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type userReponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Avater string `json:"avater"`
}

func (a *AuthController) SignUp(ctx *gin.Context) {
	var form authForm
	var user models.User

	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	copier.Copy(&user, &form)

	user.Password = user.GenerateHashPassword()

	if err := a.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var serializedAuth authReponse
	copier.Copy(&serializedAuth, &user)
	ctx.JSON(http.StatusOK, gin.H{"user": serializedAuth})
}

func (a *AuthController) GetProfile(ctx *gin.Context) {
	sub, _ := ctx.Get("sub")

	user := sub.(*models.User)

	var serializedUser userReponse
	copier.Copy(&serializedUser, user)
	ctx.JSON(http.StatusOK, gin.H{"user": serializedUser})

}
