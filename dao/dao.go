package dao

import (
	"blog/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	RegisterUser(user *model.User)
	Login(username string) model.User

	// AddPost 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
	Update(id int, post interface{})
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:111111@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed db connect", err)
	}
	Mgr = &manager{db: db}
	//创建表
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})

	fmt.Println(123)
}

func (mgr *manager) RegisterUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)

	fmt.Printf("1111%#v\n", user)
	return user
}

//博客操作

func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}

func (mgr *manager) Update(id int, post interface{})  {
	//err := mgr.db.Where("id = ?", id).Updates(post).Error
	//return err == nil
	//err := mgr.db.Find(&post, "id = ?", id)
	//return err == nil
	mgr.db.Table("posts").Where("id=?",id).Updates(
		map[string]interface{}{})
}
