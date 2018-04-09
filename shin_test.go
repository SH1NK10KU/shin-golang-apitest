package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"./util"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}
var config *util.Config
var request *util.Request
var response string

func setRequest(api string, data string) error {
	request.Path = config.Get("api." + api + ".path").(string)
	request.Method = config.Get("api." + api + ".method").(string)
	if err := json.Unmarshal([]byte(data), &request.Params); err != nil {
		panic(err)
	}
	return nil
}

func iSendRequest() error {
	response = request.SendRequest()
	return nil
}

func theParameterOfResponseShouldBe(param string, value string) error {
	if !strings.Contains(response, value) {
		return fmt.Errorf("%s应该为%s", param, value)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^(.*)接口：发送数据“(.*)”$`, setRequest)
	s.Step(`^发送请求$`, iSendRequest)
	s.Step(`^响应参数“(.*)”应为：(.*)$`, theParameterOfResponseShouldBe)

	s.BeforeScenario(func(interface{}) {
		config = new(util.Config)
		config.Read("config.json")

		request = new(util.Request)
		request.Host = config.Get("host").(string)
		request.Port = config.Get("port").(string)
		request.Headers = config.Get("headers").(map[string]interface{})
	})
}
