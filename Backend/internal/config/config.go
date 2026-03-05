package config

import "github.com/oneelabed/RSSAggregator/internal/database"

type ApiConfig struct {
	DB *database.Queries // connection to DB
}
