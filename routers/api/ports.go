package api

import (
	"AdapterServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	//"reflect"
	"github.com/buger/jsonparser"
	//"github.com/json-iterator/go"
)

func ListPorts(c *gin.Context) {

	url := "http://controller:9696/v2.0/ports"
	req, _ := http.NewRequest("GET", url, nil)

	token := c.Request.Header["X-Auth-Token"][0]
	req.Header.Add("X-Auth-Token", token)

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//FormatOutPutJson(body)
	//fmt.Println("---------------->", reflect.TypeOf(resp.Status), resp.Status)

	//content, valueType, offset, err := jsoniter.Get(body, "ports", "[0]")
	content, valueType, offset, err := jsonparser.Get(body, "ports", "[0]")

	fmt.Println("content ===>>", string(content))
	fmt.Println("valueType ===>>", valueType)
	fmt.Println("offset ===>>", offset)
	fmt.Println("err ===>>", err)

	if resp.Status == "200 OK" {
		//FormatOutPutJson(body)
		//c.Header("X-Subject-Token", resp.Header["X-Subject-Token"][0])
		//ResSuccessToClient(c, string(body))
		ResSuccessToClient(c, string(body))
		return
	}

	//FormatOutPutJson(body)
	ResErrorToClientNoArg(c, util.ERROR)
}
