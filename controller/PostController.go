package controller

import (
	"gindemo/common"
	"gindemo/model"
	"gindemo/response"
	"gindemo/vo"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)



//定义方法：先定义接口，接口中写需要实现的方法
type IPostController interface {
	RestController
	PageList(ctx * gin.Context)
}
//构造方法

func NewPostController() IPostController  {
	DB := common.GetDB()
	DB.AutoMigrate(model.Post{})

	return PostController{DB: DB}
}

//定义属性：相当于controller -> service这里直接操作dao层
type PostController struct {
	DB * gorm.DB
}

func (p PostController) Create(context *gin.Context) {
	var requestPost vo.PostRequest

	if err:=context.ShouldBind(&requestPost); err != nil{
		log.Println(err.Error())
		response.Fail(context,nil,"数据验证错误，分类名称必填！")
		return
	}
	user,_ := context.Get("user")
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.DB.Create(&post).Error; err != nil{
		panic(err)
		return
	}

	response.Success(context,nil,"创建成功")
}

func (p PostController) Update(context *gin.Context) {
	var requestPost vo.PostRequest

	if err:=context.ShouldBind(&requestPost); err != nil{
		log.Println(err.Error())
		response.Fail(context,nil,"数据验证错误，分类名称必填！")
		return
	}
	postId := context.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(context,  nil,"文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户
	user, _ := context.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(context, nil,"文章不属于您，请勿非法操作" )
		return
	}

	// 更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(context, nil,"更新失败")
		return
	}

	response.Success(context, gin.H{"post": post}, "更新成功")
}

func (p PostController) Delete(context *gin.Context) {
	// 获取path中的id
	postId := context.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(context,  nil,"文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户
	user, _ := context.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(context, nil,"文章不属于您，请勿非法操作")
		return
	}

	p.DB.Delete(&post)

	response.Success(context, gin.H{"post": post}, "成功")
}

func (p PostController) Show(context *gin.Context) {
	postId := context.Params.ByName("id")

	var post model.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(context, nil,"文章不存在")
		return
	}

	response.Success(context, gin.H{"post": post}, "成功")
}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int
	p.DB.Model(model.Post{}).Count(&total)

	// 返回数据
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}

