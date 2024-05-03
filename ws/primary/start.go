package primary

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goweb-blog/ws"
	"goweb-blog/ws/go_ws"
)

// 定义映射关系
var serveMap = map[string]ws.ServeInterface{
	"serve":   &ws.Serve{},
	"GoServe": &go_ws.GoServe{},
}

func Create() ws.ServeInterface {
	_type := viper.GetString("app.serve_type")
	return serveMap[_type]
}

func OnlineUserCount() int {
	return Create().GetOnlineUserCount()
}

func OnlineRoomUserCount(roomId int) int {
	return Create().GetOnlineRoomUserCount(roomId)
}

func Start(gin *gin.Context) {
	Create().RunWs(gin)
}
