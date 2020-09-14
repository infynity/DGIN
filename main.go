package main

import (
	. "dcgin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"

	. "dcgin/config"
	"dcgin/router"
)

func main() {
	//加载配置
	InitConfig()
	//日志配置
	InitLog()
	//数据库配置
	InitDb()
	//run
	router.Init()

}

func InitLog() {

	if AppConfig.Env == "online" {
		//使用daily加入日期
		f, _ := os.Create(fmt.Sprintf("gin_%s.log", time.Now().Format("2006-01-02")))
		// 同时将日志写入文件和控制台
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

}
