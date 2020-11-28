package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



//{
//	code:20001,
//	data : xxx,
//	msg :xx
//
//}

func Response(context * gin.Context,httpstatus int ,code int , data gin.H,msg string){
	context.JSON(httpstatus,gin.H{"code":code,"data":data,"msg":msg})

}

func Success(ctx *  gin.Context,data gin.H,msg string)  {
	Response(ctx,http.StatusOK,200,data,msg)
}

func Fail(ctx * gin.Context,data gin.H,msg string)  {
	Response(ctx,http.StatusOK,400,data,msg)
}