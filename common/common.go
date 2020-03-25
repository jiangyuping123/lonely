package common

import (
	"AdapterServer/util"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormatOutPutJson(arg []byte) {

	var str bytes.Buffer
	_ = json.Indent(&str, arg, "", "    ")

	fmt.Println("=============================================\n")
	fmt.Println(str.String())
	fmt.Println("=============================================\n")
}

func ResErrorToClientNoArg(c *gin.Context, errCode uint32) {
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": errCode,
		"ErrorDesc": util.GetMsg(errCode),
	})
}

func ResErrorToClient(c *gin.Context, errCode uint32, resContent string) {
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": errCode,
		"ErrorDesc": util.GetMsg(errCode),
		"Content":   resContent,
	})
}

func ResSuccessToClientNoArg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": util.SUCCESS,
		"ErrorDesc": util.GetMsg(util.SUCCESS),
	})
}

func ResSuccessToClient(c *gin.Context, resContent string) {
	c.JSON(http.StatusOK, gin.H{
		"ErrorCode": util.SUCCESS,
		"ErrorDesc": util.GetMsg(util.SUCCESS),
		"Content":   resContent,
	})
}
