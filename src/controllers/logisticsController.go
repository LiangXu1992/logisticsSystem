package controllers


import (
	"github.com/gin-gonic/gin"
    "common"
    "models"
    "net/http"
    "fmt"
)

//type Condition struct {
//    logistics_id string
//    logistics_name string
//    logistics_pwd string
//}

//物流员默认api
func LogisticsIndex(c *gin.Context) {
	c.JSON(common.RESPONSE_STATUS_SUCCESS, gin.H{
		"message": "pong",
	})
}

//物流员登录
func LogisticsLogin(c *gin.Context) {
    var logistics models.Shop_logistics
    logistics.Logistics_name = c.PostForm("username")
    logistics.Password       = c.PostForm("password")
    logistics.Client_type    = c.PostForm("client")
    logistics.App_token      = c.Request.Header.Get("devToken")

    fmt.Println(logistics)
    //参数是否为空
    if common.IsNull([]string{logistics.Logistics_name, logistics.Password, logistics.Client_type}) {
        c.JSON(http.StatusOK, gin.H{
            "message" : "miss params, login err",
        })
        return
    }

    //验证物流员信息并更新token，现在版本设置为单用户登录，另外一台设备登录会造成上一台掉线
    //更新登录token，极光appToken
    logistics.GetLogisticsInfo()
    if logistics.Logistics_id == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "login err",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code" :200,
        "logisticsInfo": logistics,
    })
}

//物流员订单列表
func LogisticsOrder(c *gin.Context) {
    var logistics models.Shop_logistics
    var dispatch models.Shop_v_dispatch

    logistics.Token    = c.Request.Header.Get("token")
    //判断是否登录
    logistics.IsLogin()
    if (logistics.Logistics_id == 0) {
        c.JSON(http.StatusOK, gin.H{
            "message": "login first",
        })
        return
    }

    //获取订单列表
    dispatch.GetDispatchList(&logistics)
    c.JSON(http.StatusOK, gin.H{
        "message": logistics,
    })
    return
    fmt.Println("logistics order")
    c.JSON(http.StatusOK, gin.H{
        "message": "login err",
    })
    return
}


