package syncdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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

func SyncToken() {

	authObject := AuthJson{
		Auth: Auth{
			Identity: Identity{
				Methods: []string{"password"},
				Password: Password{
					User: User{
						Name:     "jiangyuping",
						Password: "123456",
						Domain: Domain{
							Name: "Default",
						},
					},
				},
			},
		},
	}

	authJson, err := json.Marshal(authObject)
	if err != nil {
		fmt.Println("授权post参数不合法:", err)
		return
	}

	url := fmt.Sprintf("http://%s:%d/v3/auth/tokens", "controller", 5000)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(authJson)))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("授权post参数不合法:", err)
		return
	}

	if resp.Status == "201 CREATED" {

		LASTEST_TOKEN = resp.Header["X-Subject-Token"][0]

		if LASTEST_TOKEN != "" {
			TOKEN_IS_EXPIRES = false
			//fmt.Println("后台同步服务从云平台获取了token: ", LASTEST_TOKEN)
		}

	} else {
		fmt.Println("同步云平台token失败: ", resp.Status)
	}

	//body, _ := ioutil.ReadAll(resp.Body)
}
