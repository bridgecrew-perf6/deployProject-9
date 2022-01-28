package inmemory

import (
	"github.com/v-lozhkin/deployProject/internal/pkg/models"
	"github.com/v-lozhkin/deployProject/internal/pkg/user"
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
