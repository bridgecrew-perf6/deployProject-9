package user

import (
	"github.com/v-lozhkin/deployProject/internal/pkg/models"

	"github.com/labstack/echo/v4"
)

type Delivery interface {
	Login(ectx echo.Context) error
}

type Usecase interface {
	Validate(login, password string) bool
}

type Repository interface {
	List() []models.User
}
