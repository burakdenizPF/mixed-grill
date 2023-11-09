package mysql

import (
	"context"
	"fmt"
	"mixed-grill/m/v2/internal/entity"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PropertyRepo struct {
	executor sqlx.ExtContext
}

func NewPropertyRepo(executor sqlx.ExtContext) PropertyRepo {
	return PropertyRepo{executor: executor}
}

func (repo PropertyRepo) All(ctx context.Context, filter entity.PropertyFilter) ([]entity.Property, error) {
	query, args, err := repo.selectProperty(filter)
	if err != nil {
		return nil, fmt.Errorf("error building query for select property %w", err)
	}

	properties := make([]entity.Property, 0)

	err = sqlx.SelectContext(ctx, repo.executor, &properties, query, args...)
	if err != nil {
		return nil, fmt.Errorf("execute properties: %w", err)
	}

	return properties, nil
}

func (repo PropertyRepo) selectProperty(filter entity.PropertyFilter) (string, []any, error) {
	qb := sq.
		Select(
			"p.id",
			"p.location",
			"p.lng",
			"p.lat",
			"p.price",
		).
		From("property p")

	if filter.MaxLat != nil {
		qb = qb.Where(entity.ComparableFilterPredicate("p.lat", *filter.MaxLat))
	}

	if filter.MinLat != nil {
		qb = qb.Where(entity.ComparableFilterPredicate("p.lat", *filter.MinLat))
	}

	if filter.MaxLng != nil {
		qb = qb.Where(entity.ComparableFilterPredicate("p.lng", *filter.MaxLng))
	}

	if filter.MinLng != nil {
		qb = qb.Where(entity.ComparableFilterPredicate("p.lng", *filter.MinLng))
	}

	if filter.Price != nil {
		qb = qb.Where(entity.ComparableFilterPredicate("p.price", *filter.Price))
	}

	return qb.ToSql()
}
