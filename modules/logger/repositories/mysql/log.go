package mysql

import (
	"github.com/AdiSaripuloh/online-store/modules/logger/models"
	"github.com/AdiSaripuloh/online-store/modules/logger/repositories"
	"github.com/jinzhu/gorm"
	"sync"
)

type loggerRepository struct {
	db *gorm.DB
}

var (
	loggerRepoLock sync.Once
	loggerRepo     repositories.LoggerRepository
)

func NewLoggerRepository(db *gorm.DB) repositories.LoggerRepository {
	loggerRepoLock.Do(func() {
		loggerRepo = &loggerRepository{
			db: db,
		}
	})

	return loggerRepo
}

func (ur *loggerRepository) Create(log *models.Logs) (*models.Logs, error) {
	err := ur.db.Create(log).Error
	if err != nil {
		return nil, err
	}
	return log, nil
}
