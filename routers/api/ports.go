package api

import (
	"AdapterServer/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"reflect"
	"github.com/buger/jsonparser"
	//"github.com/json-iterator/go"
)

func ListPorts(c *gin.Context) {

	token := c.Request.Header["X-Auth-Token"][0]

	url := "http://controller:9696/v2.0/ports"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	params := req.URL.Query()
	params.Add("fields", "name")
	params.Add("fields", "network_id")
	params.Add("fields", "tenant_id")
	params.Add("fields", "device_owner")
	params.Add("fields", "mac_address")
	params.Add("fields", "fixed_ips")

	req.URL.RawQuery = params.Encode()
	req.Header.Add("X-Auth-Token", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	FormatOutPutJson(body)
	//fmt.Println("---------------->", reflect.TypeOf(resp.Status), resp.Status)

	//content, valueType, offset, err := jsoniter.Get(body, "ports", "[0]")
	content, valueType, offset, err := jsonparser.Get(body, "ports", "[0]")

	if err == nil {
		fmt.Println("content ===>>", string(content))
		fmt.Println("valueType ===>>", valueType)
		fmt.Println("offset ===>>", offset)
		fmt.Println("err ===>>", err)
	}

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
