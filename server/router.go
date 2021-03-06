package server

import (
	"os"
	"record-server/api"
	"record-server/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/me/session", api.UserLogout)

			auth.POST("tags", api.TagAdd)
			auth.GET("tags", api.TagList)

			auth.POST("records", api.RecordAdd)
			auth.DELETE("records/:id", api.RecordDelete)
			auth.GET("records", api.RecordList)
			auth.GET("records/:id", api.RecordDetail)
		}

	}
	return r
}
