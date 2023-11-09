package mysql

import (
	"fmt"
	"mixed-grill/m/v2/internal/config"
	"mixed-grill/m/v2/internal/datastore"

	"github.com/jmoiron/sqlx"
)

type DataStore struct {
	executor sqlx.ExtContext

	property datastore.PropertyRepo
}

func New(executor sqlx.ExtContext) DataStore {
	ds := DataStore{
		executor: executor,
	}
	ds.setup()

	return ds
}

func NewConnected(datastoreConfig config.DataStoreConfig) (DataStore, error) {
	exec, err := sqlx.Connect("mysql", datastoreConfig.URL)
	if err != nil {
		return DataStore{}, fmt.Errorf("failed to connect to mysql: %w", err)
	}

	exec.SetConnMaxLifetime(datastoreConfig.ConnMaxLifetime)
	exec.SetMaxOpenConns(datastoreConfig.MaxOpenConns)
	exec.SetMaxIdleConns(datastoreConfig.MaxIdleConns)

	if err = exec.Ping(); err != nil {
		return DataStore{}, fmt.Errorf("failed to ping connection to mysql: %w", err)
	}

	return New(exec), nil
}

func (d *DataStore) setup() {
	d.property = NewPropertyRepo(d.executor)
}

func (d DataStore) Property() datastore.PropertyRepo {
	return d.property
}
