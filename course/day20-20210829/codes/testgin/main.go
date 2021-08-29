package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// 输入 http request 中读取数据

// url
// body
// 		application/x-www-form-urlencoded
// 		multipart-form/form-data
// 		application/json
// header

// 输出 http resposne
// body
// 		html
// 		json
// header

func main() {
	rand.Seed(time.Now().Unix())
	router := gin.Default()

	count := 0
	router.Use(func(c *gin.Context) {
		fmt.Println("auth before")

		count += 1
		if count%2 == 0 {
			// 通过
			c.Next() //执行后续中间件
			fmt.Println("auth after")
			return
		}
		c.String(200, "error")
		c.Abort() // 终止处理器函数执行
	})

	// url => 处理器函数
	router.GET("/health/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ok":     true,
			"error":  "",
			"result": nil,
		})
	})

	router.POST("/test/post/", func(c *gin.Context) {
		c.String(200, "post")
	})

	router.DELETE("/test/delete/", func(c *gin.Context) {
		c.String(200, "delete")
	})

	router.PUT("/test/put/", func(c *gin.Context) {
		c.String(200, "put")
	})

	router.Any("/test/any/", func(c *gin.Context) {
		c.String(200, "any")
	})

	router.DELETE("/test/put/", func(c *gin.Context) {
		c.String(200, "put delete")
	})

	router.GET("/test/queryparams", func(c *gin.Context) {

		var params struct {
			ID int `form:"id"`
		}

		fmt.Println(c.GetQuery("id"))
		c.BindQuery(&params)

		fmt.Println(params)
		c.String(200, "queryparams")
	})

	router.POST("/test/form", func(c *gin.Context) {

		var params struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}

		fmt.Println(c.GetPostForm("username"))
		fmt.Println(c.GetPostForm("password"))
		c.Bind(&params)

		fmt.Println(params)
		c.String(200, "form")
	})

	router.POST("/test/json", func(c *gin.Context) {

		var params struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}

		c.BindJSON(&params)

		fmt.Println(params)
		c.String(200, "json")
	})

	router.POST("/test/head", func(c *gin.Context) {

		var params struct {
			Token string `form:"token"`
		}

		c.BindHeader(&params)
		fmt.Println(c.GetHeader("token"))

		fmt.Println(params)
		c.String(200, "head")
	})

	router.LoadHTMLGlob("views/*.html")
	router.GET("/index/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"Name": "kk",
		})
	})
	// 启动服务
	router.Run(":10000")
}
