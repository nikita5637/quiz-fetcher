package storage

import (
	"context"
	"database/sql"

	"github.com/nikita5637/quiz-fetcher/internal/config"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/storage/mysql"
)

// NewDB ...
func NewDB(ctx context.Context, driver string) (*sql.DB, error) {
	switch driver {
	case mysql.DriverName:
		return mysql.NewDB(ctx, config.GetMySQLDatabaseDSN())
	}

	return nil, nil
}
