package main

import (
	"os"
	"time"
	"fmt"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
)

func monitor() {
	ak := os.Getenv("JDCLOUD_AK")
	sk := os.Getenv("JDCLOUD_SK")
	cr := core.NewCredentials(ak, sk)
	client := client.NewMonitorClient(cr)
	client.JDCloudClient.Config.SetEndpoint("monitor.cn-north-1.jdcloud-api.com")

	//unit := "xx"

	a1 := models.MetricDataCm{
		Namespace: "computers",
		Metric: "cpu.load12",
		Dimensions: map[string]string{"host": "lidaobing-T470"},
		Timestamp: time.Now().Unix(),
		Type: 1,
		Values: map[string]string{"value": "xxx"},
		//Unit: &unit,
	}

	req := apis.NewPutMetricDataRequestWithAllParams([]models.MetricDataCm{a1})

	res, err := client.PutMetricData(req)

	if err != nil {
		fmt.Println("error happend")
		panic(err)
	}

	fmt.Println(res)
}

func main() {
	monitor()
}