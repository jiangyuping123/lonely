package syncdata

import (
	"AdapterServer/common"
	"fmt"
	"github.com/buger/jsonparser"
	//"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func SyncPorts() {

	fmt.Println("SyncPorts....")

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
	req.Header.Add("X-Auth-Token", LASTEST_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.Status == "200 OK" {

		common.FormatOutPutJson(body)
		content, valueType, offset, err := jsonparser.Get(body, "ports", "[0]")

		if err == nil {
			fmt.Println("content ===>>", string(content))
			fmt.Println("valueType ===>>", valueType)
			fmt.Println("offset ===>>", offset)
			fmt.Println("err ===>>", err)
		}
	}
}
