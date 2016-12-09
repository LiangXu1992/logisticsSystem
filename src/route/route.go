package route

import (
	"common"
	"controllers"
	"databases"
	"github.com/gin-gonic/gin"
)

//初始化路由表
func InitRoute(c *gin.Engine) {
	//物流根目录
	c.GET(common.LOGISTICS_ROOT, controllers.LogisticsIndex)
	//物流员登录
	c.POST(common.LOGISTICS_ROOT+"login", databases.InitDb, controllers.LogisticsLogin, databases.CloseDb)
	//物流员订单列表
	c.GET(common.LOGISTICS_ROOT+"order", databases.InitDb, controllers.LogisticsOrder, databases.CloseDb)

	//商城相关

	//商家相关

	//自己玩
	c.GET(common.LOGISTICS_ROOT+"my", controllers.My)
}
