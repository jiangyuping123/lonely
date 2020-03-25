package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

///////////////////////// 授权post json 参数 ////////////////////
//															   //
//		{												       //
//			"auth": {									       //
//				"identity": {                                  //
//					"methods": [ "password" ],                 //
//					"password": {						       //
//						"user": {						       //
//							"name": "jiangyuping",		       //
//							"password": "123456",		       //
//							"domain": {					       //
//								"name": "Default"              //
//							},                                 //
//						}                                      //
//					}                                          //
//				}                                              //
//			}                                                  //
//		}                                                      //
//                                                             //
/////////////////////////////////////////////////////////////////

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

func FormatOutPutJson(arg []byte) {
	var str bytes.Buffer
	_ = json.Indent(&str, arg, "", "    ")
	fmt.Println("formated:\n", str.String())
}

func GetTokenFromCloudPlatform(userName, password, domainName, ipAddress string, port uint16) string {

	url := fmt.Sprintf("http://%s:%d/v3/auth/tokens", ipAddress, port)

	authObject := AuthJson{
		Auth: Auth{
			Identity: Identity{
				Methods: []string{"password"},
				Password: Password{
					User: User{
						Name:     userName,
						Password: password,
						Domain: Domain{
							Name: domainName,
						},
					},
				},
			},
		},
	}

	authJson, err := json.Marshal(authObject)
	if err != nil {
		fmt.Println("授权post参数不合法:", err)
	}

	var jsonStr = []byte(authJson)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//fmt.Println("status", resp.Status)
	//fmt.Println("response:", resp.Header)
	token := resp.Header["X-Subject-Token"]
	fmt.Println(token)
	body, _ := ioutil.ReadAll(resp.Body)

	var str bytes.Buffer
	_ = json.Indent(&str, []byte(body), "", "    ")

	//fmt.Println("formated:\n", str.String())
	//return str.String()
	return ""
}

func main() {
	GetTokenFromCloudPlatform("jiangyuping", "123456", "Default", "controller", 5000)
}
