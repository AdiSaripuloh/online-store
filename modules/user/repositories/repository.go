package repositories

import (
	"github.com/AdiSaripuloh/online-store/common/responses"
	"github.com/AdiSaripuloh/online-store/modules/user/models"
)

type UserRepository interface {
	FindAll() ([]*models.User, *responses.HttpError)
}
