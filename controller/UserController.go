package controller

import (
	"gindemo/common"
	"gindemo/dto"
	"gindemo/model"
	"gindemo/response"
	"gindemo/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(context *gin.Context) {

	db := common.GetDB()

	var user = model.User{}

	context.BindJSON(&user)

	name := user.Name
	telephone := user.Telephone
	password := user.Password

	
	if len(telephone) < 11{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"电话少于11位")
		return
	}

	if len(password) < 6{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"密码少于6位")
		return
	}



	if len(name) == 0{
		name = utils.RandomString(10)
	}



	log.Println(name,telephone,password)


	if isTelephoneExist(db,telephone){
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"用户已经存在")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusUnprocessableEntity, 500, nil, "加密失败")
		return
	}

	newUser := model.User{

		Name: name,
		Telephone: telephone,
		Password: string(hashPassword),
	}

	db.Create(&newUser)
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(context, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	response.Success(context, gin.H{"token": token}, "注册成功")
}

func Login(context * gin.Context)  {

	db := common.GetDB()
	var user = model.User{}

	context.BindJSON(&user)

	telephone := user.Telephone
	password := user.Password


	if len(telephone) < 11{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"电话少于11位")
		return
	}

	if len(password) < 6{
		response.Response(context,http.StatusUnprocessableEntity,422,nil,"密码少于6位")
		return
	}


	db.Where("telephone=?",telephone).First(&user)

	if user.ID == 0{
		response.Response(context,http.StatusUnprocessableEntity,400,nil,"用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(context, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	token,err := common.ReleaseToken(user)

	if err != nil{
		response.Response(context,http.StatusUnprocessableEntity,500,nil,"系统异常")
		return
	}

	response.Response(context,http.StatusOK,200,gin.H{"token":token},"登录成功")




}

func Info(ctx *gin.Context)  {
	user,_ := ctx.Get("user")

	ctx.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{
			"user": dto.ToUserDto(user.(model.User))},
	})

}


func isTelephoneExist(db *gorm.DB , telephone string) bool{
	var user model.User
	db.Where("telephone = ?",telephone).Find(&user)

	if user.ID != 0{
		return true
	}
	return false
}
