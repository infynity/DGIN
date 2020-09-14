package middelware

import (
	"dcgin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
check if a user is a lawyer
simulating params with lid  is a lawyer
*/
func AuthLawyer() gin.HandlerFunc {
	return func(context *gin.Context) {
		lawyerId, _ := strconv.Atoi(context.Query("lawyer_id"))

		if lawyerId == 0 {
			context.JSON(http.StatusForbidden, model.Fail("you are not a lawyer"))
			context.Abort()
			return
		}

		context.Set("lawyerid", lawyerId)
		context.Next()
	}
}

/**
check if a lawyer is forbidden
simulating lid = 2379 locked
*/
func CheckForbidden() gin.HandlerFunc {
	return func(context *gin.Context) {

		lawyerId, _ := context.Get("lawyerid")
		//lawyerId,_ := strconv.Atoi(context.Query("lid"))

		if lawyerId == 2379 {
			context.JSON(http.StatusForbidden, model.Fail("you are locked,can not reply"))
			context.Abort()
			return
		}
		context.Next()
	}
}
