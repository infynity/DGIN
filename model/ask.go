package model

import "time"

//咨询模型
type Ask struct {
	Id        int       `gorm:"primary_key"`
	Content   string    `json:"content" form:"content"  binding:"required"`
	Replys    []Reply   `gorm:"foreignKey:ConsultId"`
	CreatedAt time.Time `gorm:"type:datetime;index:created_at" json:"created_at"`
}
