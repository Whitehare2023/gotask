package main

import (
	"gotask/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func main() {
	// 创建gin对象
	r := gin.Default()
	r.Use(ginsession.New())
	// 设置静态页面
	r.StaticFile("/", "./dist/index.html")
	r.StaticFile("/index", "./dist/index.html")
	r.Static("/static", "./dist/static") // 把/static路径切换为./dist/static

	// 设置路由规则
	/*
		路由规则三要素：
			1.方法：GET
			2.URL：http://localhost:9090/ping
			3.服务函数：func(c *gin.Context)
	*/

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 注册规则 curl -X POST -H "Content-type:application/json" -d '{"username":"Anna","password":"123"}' http://localhost:9090/register
	r.POST("/register", routes.Register)
	// 登录规则 curl -c anna.cookie -X POST -H "Content-type:application/json" -d '{"username":"Anna","password":"123"}' http://localhost:9090/login
	r.POST("/login", routes.Login)
	// 任务发布规则 curl -b anna.cookie -X POST -H "Content-type:application/json" -d '{"task_name":"xiwan","bonus":200}' http://localhost:9090/issue
	r.POST("/issue", routes.Issue)
	// 挖矿规则 curl -b anna.cookie -X POST -H "Content-type:application/json" -d '{"To":"Anna","Value":10000}' http://localhost:9090/mint
	r.POST("/mint", routes.Mint)
	// 任务修改规则
	// 以接受为例(status=1) curl -b grace.cookie -X POST -H "Content-type:application/json" -d '{"task_id":"0","task_status":1}' http://localhost:9090/update
	// 提交任务 curl -b grace.cookie -X POST -H "Content-type:application/json" -d '{"task_id":"0","task_status":2}' http://localhost:9090/update
	// 必须先创建另一个角色来调用 curl -X POST -H "Content-type:application/json" -d '{"username":"Grace","password":"456"}' http://localhost:9090/register
	r.POST("/update", routes.Update)
	// 任务查询 curl http://localhost:9090/tasklist?page=1
	r.GET("/tasklist", routes.Tasklist)
	// 启动web服务
	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
