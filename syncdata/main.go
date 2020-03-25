package syncdata

import (
	"fmt"
	"time"
)

var LASTEST_TOKEN string
var TOKEN_IS_EXPIRES bool

func SyncData() {

	LASTEST_TOKEN = ""
	TOKEN_IS_EXPIRES = true

	for {

		fmt.Println("============= starting sync data from cloud platform ============")

		for TOKEN_IS_EXPIRES {
			SyncToken()
			time.Sleep(time.Duration(2) * time.Second)
		}

		SyncFloatingips()
		SyncHypervisors()
		SyncLoadbalancers()
		SyncNetwork()
		SyncPorts()
		SyncProjects()
		SyncRouters()
		SyncSegments()
		SyncServers()

		time.Sleep(time.Duration(2) * time.Second)
	}
}
