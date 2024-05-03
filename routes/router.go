package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goweb-blog/controllers"
	"goweb-blog/services/session"
	"goweb-blog/static"
	"goweb-blog/ws/primary"
	"net/http"
)

func InitRoute() *gin.Engine {
	//router := gin.Default()
	router := gin.New()

	if viper.GetString(`app.debug_mod`) == "false" {
		//live 模式 打包用
		router.StaticFS("/static", http.FS(static.EmbedStatic))
	} else {
		//dev 开发用 避免修改静态资源需要重启服务
		router.StaticFS("/static", http.Dir("static"))
	}
	sr := router.Group("/", session.EnableCookieSession())
	{
		sr.GET("/", controllers.Index)
		sr.POST("/login", controllers.Login)
		sr.GET("/logout", controllers.Logout)
		sr.GET("/ws", primary.Start)

		authorized := sr.Group("/", session.AuthSessionMiddle())
		{
			authorized.GET("/home", controllers.Home)
			authorized.GET("/room/:room_id", controllers.Room)
			authorized.GET("/private-chat", controllers.PrivateChat)
			authorized.POST("/img-kr-upload", controllers.ImgkrUpload)
			authorized.GET("/pagination", controllers.Pagination)

		}
	}

	return router

}
