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

	go upHook()

	err := api.InitHttp()
	if err != nil {
		log.Fatal(err)
	}

	downSignal := make(chan os.Signal)
	signal.Notify(
		downSignal,
		syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGQUIT,
	)
	<-downSignal
	downHook()
}

func upHook() {
	country, region, city, err := getLocation()
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]interface{}{
		"action":            "up",
		"url":               config.Env.HostAddress,
		"availableServices": vpn_manager.GetAvailableServices(),
		"secret":            auth.GenerateApiToken(),
		"version":           config.Env.Version,
		"country":           country,
		"region":            region,
		"city":              city,
	}
	jsonData, _ := json.Marshal(data)

	response, err := http.Post(
		config.Env.OperatorUrl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		time.Sleep(time.Minute)
		upHook()
	}
}

func getLocation() (country string, region string, city string, err error) {

	type ipapiResponse struct {
		City    string `json:"city"`
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
	return data.Country, data.Region, data.City, nil
}

func downHook() {
	data := map[string]string{
		"action": "down",
		"url":    config.Env.HostAddress,
	}
	jsonData, _ := json.Marshal(data)

	http.Post(
		config.Env.OperatorUrl,
		"application/json",
		bytes.NewBuffer(jsonData),
	)
}
