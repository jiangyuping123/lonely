package api

import (
	"AdapterServer/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ListPorts(c *gin.Context) {

	url := "http://controller:9696/v2.0/ports"
	req, _ := http.NewRequest("GET", url, nil)

	token := c.Request.Header["X-Auth-Token"][0]
	req.Header.Add("X-Auth-Token", token)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	FormatOutPutJson(body)

	/*
		if resp.Status == "201 CREATED" {
			//FormatOutPutJson(body)
			c.Header("X-Subject-Token", resp.Header["X-Subject-Token"][0])
			//ResSuccessToClient(c, string(body))
			ResSuccessToClientNoArg(c)
			return
		}

		//FormatOutPutJson(body)
	*/

	ResErrorToClient(c, util.ERROR, string(body))
}
