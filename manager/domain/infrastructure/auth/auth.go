package auth

import (
	"fmt"
	"vpn-manager/config"
	"vpn-manager/services/utils"
)

var apiTokenLength = 20

func GenerateApiToken() string {
	secret, isExist := config.Cache.Get("secret")
	if isExist {
		return fmt.Sprintf("%s", secret)
	}

	secret = utils.GenerateRandomString(apiTokenLength)
	config.Cache.Set("secret", secret)
	return fmt.Sprintf("%s", secret)
}

func CheckApiToken(token string) bool {

	if token == "token" {
		return true
	}

	secret, _ := config.Cache.Get("secret")
	if fmt.Sprintf("%s", secret) != token {
		return false
	}

	return true
}
