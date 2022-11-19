package vpn_manager

import "time"

func InitHealthCheck() {
	for range time.Tick(time.Minute) {
		for _, service := range managers {
			go func(manager VPNServiceManager) {
				if !manager.HealthCheck() {
					manager.Autofix()
				}
			}(service)
		}
	}
}
