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
	data := map[string]interface{}{
		"action":            "up",
		"url":               os.Getenv("HOST_URL"),
		"availableServices": vpn_manager.GetAvailableServices(),
		"secret":            auth.GenerateApiToken(),
		"version":           version(),
	}
	jsonData, _ := json.Marshal(data)

	response, err := http.Post(
		os.Getenv("OPERATOR_URL"),
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil || response.StatusCode != 200 {
		time.Sleep(time.Minute)
		upHook()
	}
}

func downHook() {
	data := map[string]string{
		"action": "down",
		"url":    os.Getenv("HOST_URL"),
	}
	jsonData, _ := json.Marshal(data)

	http.Post(
		os.Getenv("OPERATOR_URL"),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
}

func version() string {
	v, err := os.ReadFile("/VERSION")
	if err != nil {
		return "0.1"
	}

	return string(v)
}
