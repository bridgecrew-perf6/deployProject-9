package postgres

import (
	"context"
	"fmt"
	repomodels "github.com/v-lozhkin/deployProject/internal/pkg/item/repository/models"
	"github.com/v-lozhkin/deployProject/internal/pkg/models"
)

// Update - accept just full state. Partial update not accepted
func (r repository) Update(ctx context.Context, item models.Item) error {
	defer r.stat.MethodDuration.WithLabels(map[string]string{"method_name": "Update"}).Start().Stop()

	query := "UPDATE item SET name=:name, description=:description, price=:price, image_link=:image_link WHERE ID=:id"

	_, err := r.db.NamedExecContext(ctx, query, repomodels.ModelToRepoItem(item))
	if err != nil {
		return fmt.Errorf("failed to update item %d: %w", item.ID, err)
	}

	return nil
}
