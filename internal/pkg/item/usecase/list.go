package usecase

import (
	"context"
	"fmt"
	"lesson9/internal/pkg/models"
)

func (u usecase) List(ctx context.Context, filter models.ItemFilter) (models.ItemList, error) {
	defer u.stat.MethodDuration.WithLabels(map[string]string{"method_name": "List"}).Start().Stop()
	if err := filter.Validate(); err != nil {
		return nil, fmt.Errorf("item's validate failed: %w", err)
	}

	list, err := u.repo.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list items from repo: %w", err)
	}

	return list, nil
}
