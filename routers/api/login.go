package api

import (
	"AdapterServer/common"
	//"AdapterServer/util"
	//"bytes"
	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	//"net/http"
)

type ReqLoginJson struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	DomainName string `json:"domainName"`
}

func Login(c *gin.Context) {

	//var reqLoginJson ReqLoginJson

	//if err := c.BindJSON(&reqLoginJson); err != nil {
	//	ResErrorToClientNoArg(c, util.INVALID_PARAMS)
	//	return
	//}
}

func Logout(c *gin.Context) {
	common.ResSuccessToClientNoArg(c)
}
