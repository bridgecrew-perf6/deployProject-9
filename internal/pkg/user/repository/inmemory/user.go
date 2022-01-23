package inmemory

import (
	"lesson9/internal/pkg/models"
	"lesson9/internal/pkg/user"
)

type inmemory struct{}

func (i inmemory) List() []models.User {
	return []models.User{
		{
			Login:    "admin",
			Password: "password",
		},
	}
}

func New() user.Repository {
	return inmemory{}
}
