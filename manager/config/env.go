package config

import "os"

type envConfig struct {
	HttpAddress string

	OpenvpnScriptDir string
	OpenvpnDataDir   string
}

var Envs *envConfig

func initEnvs() {
	Envs = &envConfig{
		HttpAddress: os.Getenv("HTTP_ADDRESS"),

		OpenvpnScriptDir: os.Getenv("OPENVPN_SCRIPT_DIR"),
		OpenvpnDataDir:   os.Getenv("OPENVPN_PERSIST_DIR"),
	}
}
