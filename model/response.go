package model

// ResultModel 内部统一返回结构数据
type ResultModel struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"Msg"`
}

const (
	CodeFail = iota
	CodeSuccess
)

// Success 返回操作成功数据：Message 序列化JSON
func Success(data interface{}) *ResultModel {
	model := &ResultModel{
		Code: CodeSuccess,
		Msg:  data,
	}
	return model
}

// Fail 返回操作成功数据：Message 序列化JSON
func Fail(data interface{}) *ResultModel {
	model := &ResultModel{
		Code: CodeFail,
		Msg:  data,
	}
	return model
}
