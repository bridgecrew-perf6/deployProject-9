package echo

import (
	"errors"
	"lesson9/internal/pkg/item"

	"github.com/labstack/echo/v4"
)

func convertToEchoError(err error) error {
	if errors.Is(err, item.ErrItemNotFound) {
		return echo.ErrNotFound
	}

	return err
}
