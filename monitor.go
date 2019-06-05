package main

import (
	"os"
	"time"
	"fmt"
	"bytes"
	"os/exec"
	"log"
	"strconv"
	"strings"

	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/client"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/apis"
	"github.com/jdcloud-api/jdcloud-sdk-go/services/monitor/models"
	"github.com/jdcloud-api/jdcloud-sdk-go/core"
	"github.com/BurntSushi/toml"
)

type config struct {
	AK string
	SK string
	Namespace string
	Dimensions map[string]string
	Metrics []metric
}

type metric struct {
	Name string
	Type string
	Value string
}

func submitData(ak, sk, namespace string, dimensions map[string]string, values map[string]float64) (err error) {
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

func shellOut(command string) (error, string, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String(), stderr.String()
}

func monitor(c *config) {
	values := map[string]float64{}

	for _, m := range(c.Metrics) {
		err, stdout, stderr := shellOut(m.Value)
		if err != nil {
			log.Printf("run command `%s' failed:\n  err: %s\n, stdout: %s\n  stderr: %s\n",
				m.Value,
				err,
				stdout,
				stderr)
			continue
		}
		value, err := strconv.ParseFloat(strings.TrimSpace(stdout), 64)
		if err != nil {
			log.Printf("output of command `%s' is not float:\n  stdout: %s\n",
				m.Value,
				stdout)
			continue
		}
		values[m.Name] = value
	}

	err := submitData(c.AK, c.SK, c.Namespace, c.Dimensions, values)
	if err != nil {
		panic(err)
	}
}

func main() {
	var c config
	_, err := toml.DecodeFile("/home/lidaobing/.lidaobing-monitor.toml", &c)
	if err != nil {
		panic(err)
	}
	monitor(&c)
}
