package controller

import (
	"fmt"
	"gindemo/common"
	"gindemo/model"
	"gindemo/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type ICategoryController interface {
	RestController
}

//相当于Java中变量定义
type CategoryController struct {
	DB * gorm.DB

}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(&model.Category{})

	return CategoryController{DB: db}
}

func (c CategoryController) Create(context *gin.Context) {

	var requestCategory model.Category

	context.BindJSON(&requestCategory)

	fmt.Println(requestCategory)

	if requestCategory.Name == ""{
		response.Fail(context,nil,"数据验证错误，分类名称必填！")
		return

	}
	c.DB.Create(&requestCategory)

	response.Success(context,gin.H{"category":requestCategory},"")
}

func (c CategoryController) Update(context *gin.Context) {

	var requestCategory model.Category

	context.BindJSON(&requestCategory)

	fmt.Println(requestCategory)


	if requestCategory.Name == ""{
		response.Fail(context,nil,"数据验证错误，分类名称必填!")
		return

	}
	categoryId , _ := strconv.Atoi(context.Params.ByName("id"))
	fmt.Println(categoryId)
	
	var updateCategory model.Category

	if c.DB.First(&updateCategory,categoryId).RecordNotFound(){
		response.Fail(context,nil,"分类不存在！")
		return
	}

	fmt.Println(updateCategory)

	c.DB.Model(&updateCategory).Update("name",requestCategory.Name)

	response.Success(context,gin.H{"data":updateCategory},"修改成功!")

}

func (c CategoryController) Delete(context *gin.Context) {
	categoryId,_ := strconv.Atoi(context.Params.ByName("id"))

	err := c.DB.Delete(&model.Category{},categoryId).Error
	if  err != nil{
		response.Fail(context,nil,"删除失败，请重试!")
		return
	}


	response.Success(context,nil,"删除成功！")
}

func (c CategoryController) Show(context *gin.Context) {

	categoryId,_ := strconv.Atoi(context.Params.ByName("id"))

	var category model.Category

	if c.DB.First(&category,categoryId).RecordNotFound(){
		response.Fail(context,nil,"分类不存在!")
		return
	}

	response.Success(context,gin.H{"category":category},"")

}
