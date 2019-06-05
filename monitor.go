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

func submitData(namespace string, dimensions map[string]string, values map[string]float64) (err error) {
	ak := os.Getenv("JDCLOUD_AK")
	sk := os.Getenv("JDCLOUD_SK")
	cr := core.NewCredentials(ak, sk)
	client := client.NewMonitorClient(cr)
	client.JDCloudClient.Config.SetEndpoint("monitor.cn-north-1.jdcloud-api.com")

	if dimensions == nil {
		hostname, err := os.Hostname()
		if err != nil {
			return err
		}
		dimensions = map[string]string{"host": hostname}
	}

	data := []models.MetricDataCm{}
	now := time.Now().Unix()

	for k, v := range values {
		data = append(data, models.MetricDataCm{
			Namespace: namespace,
			Metric: k,
			Dimensions: dimensions,
			Timestamp: now,
			Type: 1,
			Values: map[string]string{"value": fmt.Sprintf("%f", v)},
		})
	}

	req := apis.NewPutMetricDataRequestWithAllParams(data)

	res, err := client.PutMetricData(req)

	if err != nil {
		return err
	}

	fmt.Println(res)
	return nil
}

func monitor() {
	err := submitData("computers", nil, map[string]float64{"load1-test": 1.0, "load5-test": 2.0})
	if err != nil {
		panic(err)
	}
}

func main() {
	monitor()
}
