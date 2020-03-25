package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)

func main() {

	data := []byte(`{
		"ports": [
			{
					"name": "",
					"network_id": "d5799b7c-e013-474b-ba2d-63f1fc51244f",
					"tenant_id": "",
					"device_owner": "network:router_gateway",
					"mac_address": "fa:16:3e:38:74:0b",
					"fixed_ips": [
						{
								"subnet_id": "97b64822-d0eb-4cab-9caa-19a44a01daef",
								"ip_address": "192.168.0.71"
						}
					]
			},
			{
					"name": "",
					"network_id": "8248cf07-a821-4510-a32d-e838c650400c",
					"tenant_id": "6915b40a58104e73b7b02eae4366107f",
					"device_owner": "network:dhcp",
					"mac_address": "fa:16:3e:6e:c3:74",
					"fixed_ips": [
						{
								"subnet_id": "fbf9548a-a426-4839-aa94-302cc4200967",
								"ip_address": "172.17.0.2"
						}
					]
			},
		]
	}`)

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {

		if err == nil {

			fmt.Println(jsonparser.GetString(value, "network_id"))

			fixed_ips_data, dataType, offset, err := jsonparser.Get(value, "fixed_ips")
			fmt.Println(dataType, offset, err)

			jsonparser.ArrayEach(fixed_ips_data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				if subnet_id, err := jsonparser.GetString(value, "subnet_id"); err == nil {
					fmt.Println("subnet_id: ", subnet_id)
				}
			})
		}

	}, "ports")
}
