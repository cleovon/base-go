package business

import "base-go/src/models"

func HealthStatus() models.HealthCheck {
	return models.HealthCheck{
		Status: "The app is healthy",
	}
}
