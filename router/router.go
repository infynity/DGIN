package router

import (
	"dcgin/config"
	"dcgin/controller"
	"dcgin/middelware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type router struct {
	engine      *gin.Engine
	lawyerGroup *gin.RouterGroup
}

//路由总入口
func Init() {

	r := &router{engine: gin.Default()}
	r.engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello this is consult service!")
	})
	//r.engine.Use(middelware.AuthLawyer())
	r.lawyerGroup = r.engine.Group("/lawyer", middelware.AuthLawyer(), middelware.CheckForbidden())
	{
		//回复路由
		r.lawyerGroup.POST("/reply", controller.Reply)
	}

	r.engine.Handle("GET","/ask/:id",controller.List)

	//发布路由
	r.engine.POST("/publish", controller.Publish)
	//查询路由
	r.engine.GET("/consult/:id", controller.List)

	//启动监听
	r.engine.Run(fmt.Sprintf(":%s", config.AppConfig.Port))
}
