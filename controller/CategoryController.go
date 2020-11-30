package controller

import (
	"gindemo/model"
	"gindemo/repository"
	"gindemo/response"
	"gindemo/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

//相当于Java中变量定义
type CategoryController struct {
	repository repository.CategoryRepository

}

func NewCategoryController() ICategoryController {

	repository := repository.NewCategoryRepository()

	repository.DB.AutoMigrate(&model.Category{})

	return CategoryController{repository: repository}
}

func (c CategoryController) Create(context *gin.Context) {

	var requestCategory vo.CategoryRequest

	if err:=context.ShouldBind(&requestCategory); err != nil{
		response.Fail(context,nil,"数据验证错误，分类名称必填！")
		return
	}


	category,err :=	c.repository.Create(requestCategory.Name)

	if err != nil{
		panic(err)
	}

	response.Success(context,gin.H{"category":category},"")
}

func (c CategoryController) Update(context *gin.Context) {

	var requestCategory vo.CategoryRequest

	err := context.ShouldBind(&requestCategory)
	if err != nil{
		response.Fail(context,nil,"数据验证错误，分类名称必填!")
		return
	}

	categoryId , _ := strconv.Atoi(context.Params.ByName("id"))

	updateCategory, err := c.repository.SelectById(categoryId)

	if err != nil{
		response.Fail(context,nil,"分类不存在！")
		return
	}


	category, err := c.repository.Update(*updateCategory,requestCategory.Name)

	if err != nil{
		panic(err)
	}

	response.Success(context,gin.H{"data":category},"修改成功!")

}

func (c CategoryController) Delete(context *gin.Context) {
	categoryId,_ := strconv.Atoi(context.Params.ByName("id"))

	err := c.repository.DeleteById(categoryId)
	if err != nil{
		response.Fail(context,nil,"删除失败，请重试！")
		return
	}


	response.Success(context,nil,"删除成功！")
}

func (c CategoryController) Show(context *gin.Context) {

	categoryId,_ := strconv.Atoi(context.Params.ByName("id"))

	category ,err := c.repository.SelectById(categoryId)
	if err != nil {
		response.Fail(context,nil,"分类不存在")
	}

	response.Success(context,gin.H{"category":category},"")

}
