package controller

import (
	"dcgin/model"
	"dcgin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"
)

//发布咨询
func Publish(context *gin.Context) {

	ask := model.Ask{}
	if err := context.ShouldBind(&ask); err != nil {

		log.Println(err.Error())
		context.JSON(http.StatusOK, model.Fail("请检查参数"))
		return
	}

	if ask.Content == "" {
		context.JSON(http.StatusOK, model.Fail("咨询内容必须"))
		return
	}

	if utf8.RuneCountInString(ask.Content) < 10 {
		context.JSON(http.StatusOK, model.Fail("咨询字符数必须大于10"))
		return
	}

	res := service.Publish(&ask)
	context.JSON(http.StatusOK, res)
}

//用户回复咨询
func Reply(context *gin.Context) {

	reply := model.Reply{}
	if err := context.ShouldBind(&reply); err != nil {
		log.Println(err.Error())
		context.JSON(http.StatusOK, model.Fail("请检查参数"))
		return
	}

	res := service.Reply(&reply)
	context.JSON(http.StatusOK, res)

}

//查看咨询及回复
func List(context *gin.Context) {

	consultId, _ := strconv.Atoi(context.Param("id"))

	if consultId <= 0 {
		context.JSON(http.StatusOK, model.Fail("咨询id参数错误"))
		return
	}

	res := service.List(consultId)
	context.JSON(http.StatusOK, res)

}
