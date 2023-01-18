package storage

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-fetcher/internal/config"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
)

// NewDB ...
func NewDB(ctx context.Context) (*sql.DB, error) {
	switch config.GetValue("Driver").String() {
	case config.DriverMySQL:
		return mysql.NewDB(ctx, config.DriverMySQL, config.GetDatabaseDSN())
	}

	return nil, nil
}
