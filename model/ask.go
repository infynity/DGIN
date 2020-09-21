package model

import "time"

//咨询模型
type Ask struct {
	Id        int       `gorm:"primary_key"`
	Content   string    `json:"contentout" form:"contentstr"  binding:"required"`
	Replys    []Reply
	CreatedAt time.Time `gorm:"type:datetime;index:created_at" json:"created_at"`
	Uid int ` json:"uid"`
}
//`gorm:"foreignKey:ConsultId"`

type AskLists []*Ask