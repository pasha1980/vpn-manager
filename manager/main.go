package main

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vpn-manager/config"
	"vpn-manager/domain/infrastructure/api"
	"vpn-manager/domain/infrastructure/auth"
	vpn_manager "vpn-manager/domain/vpn-manager"
)

func main() {
	config.InitConfig()

	go vpn_manager.InitHealthCheck()

	err := api.InitHttp()
	if err != nil {
		log.Fatal(err)
	}

	go upHook()

	downSignal := make(chan os.Signal)
	signal.Notify(
		downSignal,
		syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGQUIT,
	)
	<-downSignal
	downHook()
}

func upHook() {
	country, region, err := getLocation()

	data := map[string]interface{}{
		"action":            "up",
		"url":               config.Envs.HostAddress,
		"availableServices": vpn_manager.GetAvailableServices(),
		"secret":            auth.GenerateApiToken(),
		"version":           config.Envs.Version,
		"country":           country,
		"region":            region,
	}
	jsonData, _ := json.Marshal(data)

	response, err := http.Post(
		config.Envs.OperatorUrl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil || response.StatusCode != 200 {
		time.Sleep(time.Minute)
		upHook()
	}
}

func getLocation() (country string, region string, err error) {

	type ipapiResponse struct {
		Region  string `json:"region"`
		Country string `json:"country_name"`
	}

	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var data ipapiResponse

	json.NewDecoder(resp.Body).Decode(&data)
	return data.Country, data.Region, nil
}

func downHook() {
	data := map[string]string{
		"action": "down",
		"url":    config.Envs.HostAddress,
	}
	jsonData, _ := json.Marshal(data)

	http.Post(
		config.Envs.OperatorUrl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
}
