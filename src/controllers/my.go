package controllers

import (
    "github.com/gin-gonic/gin"

    "strconv"
    "common"
)

const (
    //增长率10%
    UP_RATE = 0.1
    //新股首天增长率
    NEW_UP_RATE = 0.44
)

func My(c *gin.Context) {
    //本金
    money, _ := strconv.ParseFloat(c.Query("money"), 64)

    //涨停天数
    upDay, _ := strconv.Atoi(c.Query("up_day"))

    //是否新股 1:是；0：否
    newShares := c.Query("new_share")

    //判断是否新股
    if (newShares == "1" && upDay >= 1) {
        money = money + money * NEW_UP_RATE
    } else {
        money = money + money * UP_RATE
    }

    //开始计算
    for i := 1; i < upDay; i++ {
        money = money + money * UP_RATE
    }

    c.JSON(common.RESPONSE_STATUS_SUCCESS, gin.H{
        "total is": money,
        "gujia is ": money/500,
    })

}