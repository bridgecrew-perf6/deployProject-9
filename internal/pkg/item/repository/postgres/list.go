package postgres

import (
	"context"
	"fmt"
	repomodels "github.com/v-lozhkin/deployProject/internal/pkg/item/repository/models"
	"github.com/v-lozhkin/deployProject/internal/pkg/models"
	"strings"
)

func (r repository) List(ctx context.Context, filter models.ItemFilter) (models.ItemList, error) {
	defer r.stat.MethodDuration.WithLabels(map[string]string{"method_name": "List"}).Start().Stop()

	res := repomodels.ItemList{}

	query := strings.Builder{}
	query.WriteString("SELECT * FROM item where true")
	args := make([]interface{}, 0)

	counter := 0
	if filter.ID != nil {
		counter++
		query.WriteString(fmt.Sprintf(" and id = $%d", counter))
		args = append(args, *filter.ID)
	}
	if filter.PriceMax != nil {
		counter++
		query.WriteString(fmt.Sprintf(" and price <=  $%d", counter))
		args = append(args, *filter.PriceMax)
	}
	if filter.PriceMin != nil {
		counter++
		query.WriteString(fmt.Sprintf(" and price >= $%d", counter))
		args = append(args, *filter.PriceMin)
	}

	stmt, err := r.db.PreparexContext(ctx, query.String())
	if err != nil {
		return nil, fmt.Errorf("faield to prepare statement: %w", err)
	}

	if err = stmt.SelectContext(ctx, &res, args...); err != nil {
		return nil, fmt.Errorf("failed to select item from db: %w", err)
	}

	return repomodels.RepoItemListToModel(res), nil
}
