package route

import (
	"controllers"
	"github.com/gin-gonic/gin"
	"common"
)



//初始化路由表
func InitRoute(c *gin.Engine) {
	//物流根目录
	c.GET(common.LOGISTICS_ROOT, controllers.LogisticsIndex)
    //物流员登录
    c.POST(common.LOGISTICS_ROOT + "login", controllers.LogisticsLogin)
	//物流员订单列表
	c.GET(common.LOGISTICS_ROOT + "order", controllers.LogisticsOrder)

	//商城相关

	//商家相关

	//自己玩
    c.GET(common.LOGISTICS_ROOT + "my", controllers.My)
}
