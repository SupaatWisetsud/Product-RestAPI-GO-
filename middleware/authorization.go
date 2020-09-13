package middleware

import (
	"net/http"
	"product/models"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {
	sub, ok := ctx.Get("sub")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user := sub.(*models.User)

	enforcer := casbin.NewEnforcer("config/acl_model.conf", "config/policy.csv")
	ok = enforcer.Enforce(user, ctx.Request.URL.Path, ctx.Request.Method)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you are not allowed to access this resource"})
		return
	}

	ctx.Next()
}
