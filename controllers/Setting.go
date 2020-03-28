package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "www.alisleepy.com/models"
	//"fmt"
)

//GetMyInfo 获取个人信息和浏览次数
func GetMyInfo(this *gin.Context) {
	myInfo := GetMyInfosAndViewNum()
	this.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": myInfo,
	})
}

//GetFriendlyUrl 获取友链
func GetFriendlyUrl(this *gin.Context) {
	urls := GetFriendlyUrls()
	this.JSON(http.StatusOK, gin.H{
		"data": urls,
	})
}
