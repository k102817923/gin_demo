package main

import (
	"fmt"
	"net/http"

	"gin_demo/pkg/setting"
	"gin_demo/routers"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 	"message": "pong",
	// 	})
	// })
	// // listen and serve on 0.0.0.0:8080
	// r.Run()

	// // 返回Gin的type Engine struct{...}, 里面包含RouterGroup, 相当于创建一个路由Handlers, 可以后期绑定各类的路由规则和函数、中间件等
	// router := gin.Default()
	// // 创建不同的HTTP方法绑定到Handlers中, 也支持POST、PUT、DELETE、PATCH、OPTIONS、HEAD等常用的Restful方法
	// // Context是gin中的上下文, 它允许我们在中间件之间传递变量、管理流、验证JSON请求、响应JSON请求等, 在gin中包含大量Context的方法, 例如我们常用的DefaultQuery、Query、DefaultPostForm、PostForm等等
	// router.GET("/test", func(c *gin.Context) {
	// // 就是一个map[string]interface{}
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// // http.ListenAndServe和r.Run()没有本质区别
	// s.ListenAndServe()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
