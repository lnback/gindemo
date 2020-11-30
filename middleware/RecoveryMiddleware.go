package middleware

import (
	"fmt"
	"gindemo/response"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc  {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil{

				response.Fail(context,nil,fmt.Sprint(err))
			}

		}()

		context.Next()
	}
}
