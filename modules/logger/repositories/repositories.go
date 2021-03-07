package repositories

import "github.com/AdiSaripuloh/online-store/modules/logger/models"

type LoggerRepository interface {
	Create(log *models.Logs) (*models.Logs, error)
}
