package api

import (
	"AdapterServer/util"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	//"reflect"
)

///////////////////////// 授权post json 参数 /////////
//													//
//	{												//
//		"auth": {									//
//			"identity": {                           //
//				"methods": [ "password" ],          //
//				"password": {						//
//					"user": {						//
//						"name": "jiangyuping",		//
//						"password": "123456",		//
//						"domain": {					//
//							"name": "Default"       //
//						},                          //
//					}                               //
//				}                                   //
//			}                                       //
//		}                                           //
//	}                                               //
//                                                  //
//////////////////////////////////////////////////////

type Domain struct {
	Name string `json:"name"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   `json:"domain"`
}

type Password struct {
	User `json:"user"`
}

type Identity struct {
	Methods  []string `json:"methods"`
	Password `json:"password"`
}

type Auth struct {
	Identity `json:"identity"`
}

type AuthJson struct {
	Auth `json:"auth"`
}

type ReqLoginJson struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	DomainName string `json:"domainName"`
}

func Login(c *gin.Context) {

	var reqLoginJson ReqLoginJson

	if err := c.BindJSON(&reqLoginJson); err != nil {
		ResErrorToClientNoArg(c, util.INVALID_PARAMS)
		return
	}

	authObject := AuthJson{
		Auth: Auth{
			Identity: Identity{
				Methods: []string{"password"},
				Password: Password{
					User: User{
						Name:     reqLoginJson.UserName,
						Password: reqLoginJson.Password,
						Domain: Domain{
							Name: reqLoginJson.DomainName,
						},
					},
				},
			},
		},
	}

	authJson, err := json.Marshal(authObject)
	if err != nil {
		fmt.Println("授权post参数不合法:", err)
		ResErrorToClientNoArg(c, util.INVALID_PARAMS)
		return
	}

	var jsonStr = []byte(authJson)
	url := fmt.Sprintf("http://%s:%d/v3/auth/tokens", "controller", 5000)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("===>>> 授权post参数不合法:", err)
		ResErrorToClientNoArg(c, util.ERROR)
		return
	}

	//fmt.Println(resp.Status)
	//fmt.Println(reflect.TypeOf(resp))
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.Status == "201 CREATED" {
		//FormatOutPutJson(body)
		c.Header("X-Subject-Token", resp.Header["X-Subject-Token"][0])
		//ResSuccessToClient(c, string(body))
		ResSuccessToClientNoArg(c)
		return
	}

	//FormatOutPutJson(body)
	ResErrorToClient(c, util.ERROR, string(body))
}

func Logout(c *gin.Context) {
	ResSuccessToClientNoArg(c)
}
