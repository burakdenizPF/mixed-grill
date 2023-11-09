package config

import "time"

type (
	PortsConfig struct {
		HTTP uint `mapstructure:"http"`
	}

	DataStoreConfig struct {
		URL             string        `mapstructure:"url"`
		ConnMaxLifetime time.Duration `mapstructure:"conn_max_life_time"`
		MaxOpenConns    int           `mapstructure:"max_open_conns"`
		MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	}

	Config struct {
		Ports     PortsConfig     `mapstructure:"ports"`
		DataStore DataStoreConfig `mapstructure:"data_store"`
	}
)
