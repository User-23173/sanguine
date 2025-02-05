// Package sql provides a sql store for the screener-api.
package sql

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/contrib/screener-api/db"
	"github.com/synapsecns/sanguine/contrib/screener-api/db/sql/mysql"
	"github.com/synapsecns/sanguine/contrib/screener-api/db/sql/sqlite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
)

// Connect connects to the database.
func Connect(ctx context.Context, dbType dbcommon.DBType, path string, metrics metrics.Handler) (db.RuleDB, error) {
	switch dbType {
	case dbcommon.Mysql:
		store, err := mysql.NewMysqlStore(ctx, path, metrics)
		if err != nil {
			return nil, fmt.Errorf("could not create mysql store: %w", err)
		}

		return store, nil
	case dbcommon.Sqlite:
		store, err := sqlite.NewSqliteStore(ctx, path, metrics, false)
		if err != nil {
			return nil, fmt.Errorf("could not create sqlite store: %w", err)
		}

		return store, nil
	case dbcommon.Clickhouse:
		return nil, errors.New("driver not supported")
	default:
		return nil, fmt.Errorf("unsupported driver: %s", dbType)
	}
}
