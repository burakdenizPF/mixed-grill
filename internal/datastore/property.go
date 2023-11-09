package datastore

import (
	"context"
	"mixed-grill/m/v2/internal/entity"
)

type PropertyRepo interface {
	All(ctx context.Context, filter entity.PropertyFilter) ([]entity.Property, error)
}
