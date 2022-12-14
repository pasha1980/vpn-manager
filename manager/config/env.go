package config

import (
	"os"
)

type envConfig struct {
	HttpAddress string
	HostAddress string
	OperatorUrl string
	Version     string

	OpenvpnScriptDir string
	OpenvpnDataDir   string
}

var Env *envConfig

func initEnvs() {
	var version string
	v, err := os.ReadFile("/VERSION")
	if err != nil {
		version = "0.1"
	}

	version = string(v)

	Env = &envConfig{
		HttpAddress: os.Getenv("HTTP_ADDRESS"),
		HostAddress: os.Getenv("HOST_URL"), // todo: if ends on "/"
		OperatorUrl: os.Getenv("OPERATOR_URL"),
		Version:     version,

		OpenvpnScriptDir: os.Getenv("OPENVPN_SCRIPT_DIR"),
		OpenvpnDataDir:   os.Getenv("OPENVPN_PERSIST_DIR"),
	}
}
