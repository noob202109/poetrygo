package main

import (
	"PoetryWeb/data"
	"PoetryWeb/model"
	"PoetryWeb/utils/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 创建一个默认的路由引擎
	r := gin.Default() //带log等中间件
	// r := gin.New()//不带中间件的路由引擎

	//加载静态文件
	r.Static("/static", "./static")                      //多个文件(整个文件夹)
	r.StaticFile("/favicon.ico", "./static/favicon.ico") //单个文件

	//加载模板文件
	r.LoadHTMLGlob("view/*")

	//相应路由规则 执行的函数
	r.GET("/", Index)
	r.GET("/epoch", Epoch)
	r.GET("/poet", Poet)
	r.GET("/search", Search)
	r.GET("/content", Content)

	// 启动HTTP服务，默认在127.0.0.1:8080启动服务
	r.Run(":" + utils.Conf.Port) //也可以写全乎了比如"127.0.0.1:80"或者域名

}

//Index 主页路由规则方法
func Index(c *gin.Context) {
	//接收get参数 没有值给一个默认值
	page := c.DefaultQuery("page", "1") //获取页码

	//初始化结构体
	var index data.Index
	//调用获取导航栏目方法
	navbar := model.Getnavbar()
	//调用获取分页及分页数据方法
	pagination, poetry := model.Getpageindexs(page)
	//给相应字段赋值
	index.Navbar = navbar
	index.Page = pagination
	index.Poetry = poetry
	//输出网页
	c.HTML(http.StatusOK, "index.html", index)
}

//Epoch 朝代路由规则方法
func Epoch(c *gin.Context) {
	//接收get参数 没有值给一个默认值
	page := c.DefaultQuery("page", "1")  //获取页码
	epoch := c.DefaultQuery("epoch", "") //获取朝代

	//初始化结构体
	var other data.Other
	//调用获取导航栏目方法
	navbar := model.Getnavbar()
	//调用获取分页及分页数据方法
	pagination, poetry := model.Getpageepochs(page, epoch)
	//给相应字段赋值
	other.Navbar = navbar
	other.Page = pagination
	other.Poetry = poetry
	other.Param = epoch
	//输出网页
	c.HTML(http.StatusOK, "epoch.html", other)
}

//Poet 诗人路由规则方法
func Poet(c *gin.Context) {
	//接收get参数 没有值给一个默认值
	page := c.DefaultQuery("page", "1") //获取页码
	poet := c.DefaultQuery("poet", "")  //获取朝代

	//初始化结构体
	var other data.Other
	//调用获取导航栏目方法
	navbar := model.Getnavbar()
	//调用获取分页及分页数据方法
	pagination, poetry := model.Getpagepoets(page, poet)
	//给相应字段赋值
	other.Navbar = navbar
	other.Page = pagination
	other.Poetry = poetry
	other.Param = poet
	//输出网页
	c.HTML(http.StatusOK, "poet.html", other)
}

//Search 搜索路由规则方法
func Search(c *gin.Context) {
	//接收get参数 没有值给一个默认值
	page := c.DefaultQuery("page", "1")      //获取页码
	keyword := c.DefaultQuery("keyword", "") //获取朝代

	//初始化结构体
	var other data.Other
	//调用获取导航栏目方法
	navbar := model.Getnavbar()
	//调用获取分页及分页数据方法
	pagination, poetry := model.Getpagesearchs(page, keyword)
	//给相应字段赋值
	other.Navbar = navbar
	other.Page = pagination
	other.Poetry = poetry
	other.Param = keyword
	//输出网页
	c.HTML(http.StatusOK, "search.html", other)
}

//Content 诗词路由规则方法
func Content(c *gin.Context) {
	//接收get参数 没有值给一个默认值
	id := c.DefaultQuery("id", "1") //获取朝代

	//初始化结构体
	var content data.Content
	//调用获取导航栏目方法
	navbar := model.Getnavbar()
	//调用获取诗词数据方法
	poetry, poetry1 := model.Getcontent(id)
	//给相应字段赋值
	content.Navbar = navbar
	content.Poetry = poetry
	content.Poetry1 = poetry1
	//输出网页
	c.HTML(http.StatusOK, "content.html", content)
}
