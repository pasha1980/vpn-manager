package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"vpn-manager/config"
	"vpn-manager/domain/infrastructure/api"
	"vpn-manager/domain/infrastructure/auth"
)

func main() {
	config.InitConfig()

	go upHook()

	api.InitHttp()

	downSignal := make(chan os.Signal)
	signal.Notify(
		downSignal,
		syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGABRT, syscall.SIGQUIT,
	)
	<-downSignal
	downHook()
}

func upHook() {
	data := map[string]string{
		"action":            "up",
		"url":               os.Getenv("HOST_URL"),
		"availableServices": "openvpn", // todo
		"secret":            auth.GenerateApiToken(),
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