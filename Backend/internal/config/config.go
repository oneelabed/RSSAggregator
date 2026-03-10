package config

import "github.com/oneelabed/IsraelConflictMonitor/internal/database"

type ApiConfig struct {
	DB *database.Queries // connection to DB
}
