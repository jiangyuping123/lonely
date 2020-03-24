package main

import (
	"fmt"
	"github.com/json-iterator/go" // 引入
	"time"
)

func syncCloudData() {

	for {
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("=============>>>AAA...")
	}
}

func main() {
	val := []byte(`{"ID":1,"Name":"Reds","Colors": {"c":"Crimson","r":"Red","rb":"Ruby","m":"Maroon","tests":["tests_1","tests_2","tests_3","tests_4"]}}`)

	//fmt.Println(jsoniter.Get(val, "Colors", "tests").ToString())
	fmt.Println(jsoniter.Get(val, "Colors", "tests").Size())

	go syncCloudData()

	for {
		time.Sleep(time.Duration(10) * time.Second)
		fmt.Println("=============>>>BBB...")
	}
}
