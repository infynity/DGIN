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

	if err := model.GetDB().Where("id=?", reply.AskId).First(&ask).Error; err != nil {
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
	//ask := model.Askoutput{}
	//预加载模式一对多关联
	//
	//var adminAreaStructs []model.Reply
	//model.GetDB().Model(&model.Reply{}).Where("id in (?)",[]int{1,2}).Scan(&adminAreaStructs)
	//model.GetDB().Model(&model.Reply{}).Where("id in (?)",[]int{1,2}).Find(&adminAreaStructs)   //双次前缀
	//fmt.Println(adminAreaStructs)


	//adm:=model.Ask{Id: 1,Content: "bbbbb",Uid: 123}
	//sd:= model.GetDB().Select("id","content","uid").Save(&adm).Error
	//fmt.Println(sd,22222)
	model.GetDB().Model(model.Ask{}).Preload("Replys").First(&ask, id)
	//model.GetDB().Model(model.Ask{}).Select("id,content,created_at").Where("id=?",id).Preload("Replys").Scan(&ask)

	//err := model.GetDB().Create(&model.Ask{Content: "1234"}).Error
	//fmt.Println(err)
	//model.Logger.Println(err)
	if ask.Id <= 0 {
		return model.Fail("没有查询到该咨询信息")
	}


	return model.Success(ask)
}
