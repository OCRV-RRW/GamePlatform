package dbconn

import (
	"context"
	"fmt"
	"gameplatform/internal/database"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConnection struct {
	Pool    *pgxpool.Pool
	Queries *database.Queries
}

func NewDatabaseConnection(databaseURL string, ctx context.Context) DatabaseConnection {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't load config: %s", err.Error()))
		os.Exit(1)
	} else {
		slog.Info("✅ Connected Successfully to the Database")
	}

	queries := database.New(pool)

	return DatabaseConnection{
		Pool:    pool,
		Queries: queries,
	}
}

func (conn DatabaseConnection) CloseConnection(ctx context.Context) {
	// conn.Connection.Close(ctx)
}
