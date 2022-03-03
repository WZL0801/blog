package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"html/template"

	"strconv"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}
	//执行数据库添加
	dao.Mgr.RegisterUser(&user)

	c.Redirect(301, "/")

}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("123456", username)
	//fmt.Println("qqqq", password)
	u := dao.Mgr.Login(username)
	fmt.Println(u.Username)
	fmt.Println(u.Password)

	if u.Username == " " {
		c.HTML(200, "login.html", "用户名不存在")
		fmt.Println("用户名不存在")
	} else {

		if u.Password != password {
			fmt.Println("密码错误")
			c.HTML(200, "login.html", "密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}

}

// GetPostIndex 博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

// AddPost 添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	context := c.PostForm("context")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: context,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

// GoAddPost 跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

// PostDetail 查看详细博客
func PostDetail(c *gin.Context) {
	//获取传递的id
	s := c.Query("id")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)
	content := blackfriday.Run([]byte(p.Content))
	c.HTML(200, "detail.html", gin.H{
		"Pid":     pid,
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}

//UpDate 修改博客
func UpDate(c *gin.Context) {
	s := c.Query("id")
	pid, _ := strconv.Atoi(s)
	fmt.Println(pid)

	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("context")

	post := model.Post{

		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.Update(pid,post)
	c.HTML(200, "/post_index", gin.H{
		"Title":   title,
		"Content": content,
		"Tag":     tag,
	})

}

// GoUpdate 跳转到修改博客
func GoUpdate(c *gin.Context) {
	//获取传递的id
	s := c.Query("id")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)


	c.HTML(200, "updateIndex.html", gin.H{
		"Pid":     p.ID,
		"Title":   p.Title,
		"Content": p.Content,
		"Tag":     p.Tag,
	})

}
