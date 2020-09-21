package service

type AskService struct {

}
//
//func (this *AskService)LoadAsklists(req ctx.asklistRequest)(*model.Ask,error){
//
//	asks := &model.Ask{}
//	return asks,nil
//}

func newAskService()*AskService{
	return &AskService{}
}
