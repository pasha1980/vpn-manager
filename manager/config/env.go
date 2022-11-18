package config

import "os"

type envConfig struct {
	HttpAddress      string
	OpenvpnScriptDir string
}

var Envs *envConfig

func initEnvs() {
	Envs.HttpAddress = os.Getenv("HTTP_ADDRESS")
	Envs.OpenvpnScriptDir = os.Getenv("OPENVPN_SCRIPT_DIR")
}
