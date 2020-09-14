package service

import (
	"dcgin/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

//发布服务方法
func Publish(ask *model.Ask) *model.ResultModel {

	if err := model.GetDB().Create(&ask).Error; err != nil {
		log.Println("Create ask  Create Failed:%v", err)
		return model.Fail("创建咨询失败")

	}

	return model.Success(fmt.Sprintf("发布成功，咨询id:%d", ask.Id))
}

//回复服务方法
func Reply(reply *model.Reply) *model.ResultModel {

	ask := model.Ask{}

	if err := model.GetDB().Where("id=?", reply.ConsultId).First(&ask).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return model.Fail("查询失败")

		}
	}

	if ask.Id <= 0 {
		return model.Fail("没有该条源咨询")

	}
	if err := model.GetDB().Create(&reply).Error; err != nil {
		log.Println("Create reply  Create Failed:%v", err)
		return model.Fail("创建回复失败")
	}

	return model.Success(reply)

}

//查询服务方法
func List(id int) *model.ResultModel {
	ask := model.Ask{}
	//预加载模式一对多关联
	model.GetDB().Preload("Replys").First(&ask, id)
	if ask.Id <= 0 {
		return model.Fail("没有查询到该咨询信息")
	}
	return model.Success(ask)
}
