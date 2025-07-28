package dbconn

import (
	"context"
	"fmt"
	"gameplatform/internal/database"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

type DatabaseConnection struct {
	Queries    *database.Queries
	Connection *pgx.Conn
}

func NewDatabaseConnection(databaseURL string, ctx context.Context) DatabaseConnection {
	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't load config: %s", err.Error()))
		os.Exit(1)
	} else {
		slog.Info("✅ Connected Successfully to the Database")
	}

	queries := database.New(conn)

	return DatabaseConnection{
		Queries:    queries,
		Connection: conn,
	}
}

func (conn DatabaseConnection) CloseConnection(ctx context.Context) {
	conn.Connection.Close(ctx)
}
