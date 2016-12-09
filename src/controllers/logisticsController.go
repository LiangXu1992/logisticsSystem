package controllers

import (
	"common"
	"fmt"
	"github.com/gin-gonic/gin"
	"models"
	"net/http"
)

//物流员默认api
func LogisticsIndex(c *gin.Context) {
	c.JSON(common.RESPONSE_STATUS_SUCCESS, gin.H{
		"message": "pong",
	})
}

//物流员登录
func LogisticsLogin(c *gin.Context) {
	//初始化结构体变量
	logistics := models.Shop_logistics{
		Logistics_name: c.PostForm("username"),
		Password:       c.PostForm("password"),
		Client_type:    c.PostForm("client"),
		App_token:      c.Request.Header.Get("devToken"),
		Db:             c.Db,
	}

	//参数是否为空
	if common.IsNull([]string{logistics.Logistics_name, logistics.Password, logistics.Client_type}) {
		c.JSON(http.StatusOK, gin.H{
			"message": "miss params, login err",
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
		"code":          200,
		"logisticsInfo": logistics,
	})

}

//物流员订单列表
func LogisticsOrder(c *gin.Context) {
	fmt.Println("a")
	logistics := models.Shop_logistics{
		Db: c.Db,
	}
	dispatch := models.Shop_v_dispatch{
		Db: c.Db,
	}

	logistics.Token = c.Request.Header.Get("token")
	//判断是否登录
	str, err := logistics.IsLogin()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": str,
		})
		return
	}

	//获取订单列表
	dispatchList := dispatch.GetDispatchList(&logistics)
	c.JSON(http.StatusOK, gin.H{
		"message": dispatchList,
	})
	return
	fmt.Println("logistics order")
	c.JSON(http.StatusOK, gin.H{
		"message": "login err",
	})
	return
}
