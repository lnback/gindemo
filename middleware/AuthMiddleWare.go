package middleware

import (
	"gindemo/common"
	"gindemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc  {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer"){
			context.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]

		token,claims,err := common.ParseToken(tokenString)

		if err != nil || !token.Valid{
			context.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			context.Abort()
			return
		}


		userId := claims.UserId
		db := common.GetDB()

		var user model.User

		db.First(&user,userId)

		if user.ID == 0{
			context.JSON(http.StatusUnauthorized,gin.H{
				"code":401,
				"msg":"权限不足",
			})
			context.Abort()
			return
		}

		context.Set("user",user)

		context.Next()
	}
}
