package middleware

import (
	"log"
	"net/http"
	"os"
	"product/config"
	"product/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type authForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Authentication() *jwt.GinJWTMiddleware {
	identityKey := "sub"
	db := config.GetDB()

	authenticate, err := jwt.New(&jwt.GinJWTMiddleware{
		Key: []byte(os.Getenv("SECRET")),

		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",

		IdentityHandler: func(c *gin.Context) interface{} {
			var user models.User
			claims := jwt.ExtractClaims(c)

			id := claims[identityKey]
			if db.First(&user, uint(id.(float64))).RecordNotFound() {
				return nil
			}

			return &user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var form authForm
			var user models.User

			if err := c.ShouldBindJSON(&form); err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
				return nil, jwt.ErrMissingLoginValues
			}

			if db.Where("email = ?", form.Email).First(&user).RecordNotFound() {
				return nil, jwt.ErrFailedAuthentication
			}

			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &user, nil
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return authenticate
}
