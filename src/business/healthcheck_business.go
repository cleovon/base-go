package business

import (
	"base-go/src/models"
	"base-go/src/modules/logger"
)

func HealthStatus() *models.HealthCheck {
	log := logger.New()
	hc := new(models.HealthCheck)
	hc.Status = "The app is healthy"
	log.Info("The response was: " + string(hc.ToBytes()))
	return hc
}
