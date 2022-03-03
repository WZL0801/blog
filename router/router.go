package router

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func IndexHandler() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.POST("/register", controller.RegisterUser)
	r.GET("/register", controller.GoRegister)

	r.POST("/login", controller.Login)
	r.GET("/login", controller.GoLogin)

	r.GET("/", controller.Index)

	//博客列表
	r.GET("/post_index", controller.GetPostIndex)
	//添加博客
	r.POST("/post", controller.AddPost)
	//跳转到添加博客页面
	r.GET("/post", controller.GoAddPost)

	//详细博客查询
	r.GET("/detail", controller.PostDetail)
	//修改博客
	r.POST("/updateIndex", controller.UpDate)
	//跳转到修改博客
	r.GET("/updateIndex", controller.GoUpdate)
	r.Run()
}
