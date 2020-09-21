package model

import "time"

//咨询回复模型
type Reply struct {
	Id        int
	AskId	 int       `json:"-" form:"consult_id" binding:"required"`
	Reply     string    `json:"reply" form:"reply" binding:"required"`
	LawyerId  int       `json:"lawyer_id" form:"lawyer_id" binding:"required"`
	CreatedAt time.Time `gorm:"type:datetime;index:created_at" json:"created_at"`
}
func (as *Reply) TableName() string {

	return "" + "reply"
}
