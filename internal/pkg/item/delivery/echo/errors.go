package echo

import (
	"errors"
	"github.com/v-lozhkin/deployProject/internal/pkg/item"

	"github.com/labstack/echo/v4"
)

func convertToEchoError(err error) error {
	if errors.Is(err, item.ErrItemNotFound) {
		return echo.ErrNotFound
	}

	return err
}
