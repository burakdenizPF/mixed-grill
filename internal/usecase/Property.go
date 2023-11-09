package usecase

import (
	"context"
	"fmt"
	"mixed-grill/m/v2/internal/datastore"
	"mixed-grill/m/v2/internal/entity"
)

type Manager struct {
	ds datastore.Datastore
}

func NewManager(ds datastore.Datastore) Manager {
	return Manager{
		ds: ds,
	}
}

func (m Manager) Property(ctx context.Context, filter entity.PropertyFilter) ([]entity.Property, error) {
	// long lat calculate

	properties, err := m.ds.Property().All(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrive properties from db %w", err)
	}

	return properties, nil
}
