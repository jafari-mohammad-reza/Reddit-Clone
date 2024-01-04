package postgres_repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)

type PostgresBaseRepository[T any] struct {
    cfg   *config.Config
    lg    custome_logger.Logger
    Db    *sql.DB
    table string
}

func NewPostgresBaseRepository[T any](cfg *config.Config, lg *custome_logger.Logger, tableName string) *PostgresBaseRepository[T] {
    var repoCfg *config.Config
    var repoLg custome_logger.Logger

    if cfg == nil {
        repoCfg = config.GetConfig()
    } else {
        repoCfg = cfg
    }

    if lg == nil {
        repoLg = custome_logger.NewLogger(cfg)
    } else {
        repoLg = *lg
    }

    return &PostgresBaseRepository[T]{
        cfg:   repoCfg,
        lg:    repoLg,
        Db:    postgres.GetPostgres(),
        table: tableName,
    }
}

type BaseRepository[T any] interface {
    Create(ctx context.Context, entity any) error
    
    Delete(ctx context.Context, id string) error
    GetAll(ctx context.Context) ([]T, error)
    GetById(ctx context.Context, id string) (T, error)
    Drop(ctx context.Context) error
}

// Create inserts a new entity into the database.
func (r *PostgresBaseRepository[T]) Create(ctx context.Context, entity T) error {
    jsonData, err := json.Marshal(entity)
    if err != nil {
        return err
    }

    var columns []string
    var placeholders []string
    var args []interface{}

    // Assuming the entity is a map after JSON unmarshalling
    dataMap := make(map[string]interface{})
    json.Unmarshal(jsonData, &dataMap)

    i := 1
    for k, v := range dataMap {
        columns = append(columns, k)
        placeholders = append(placeholders, fmt.Sprintf("$%d", i))
        args = append(args, v)
        i++
    }

    query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
        r.table, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

    _, err = r.Db.ExecContext(ctx, query, args...)
    return err
}

// Delete removes an entity from the database.
func (r *PostgresBaseRepository[T]) Delete(ctx context.Context, id string) error {
    query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", r.table)
    _, err := r.Db.ExecContext(ctx, query, id)
    return err
}

// GetAll retrieves all entities from the database.
func (r *PostgresBaseRepository[T]) GetAll(ctx context.Context) ([]T, error) {
    query := fmt.Sprintf("SELECT * FROM %s", r.table)
    rows, err := r.Db.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var entities []T
    for rows.Next() {
        var entity T
        // Scan data into entity, depends on the structure of T
        if err := rows.Scan(&entity); err != nil {
            return nil, err
        }
        entities = append(entities, entity)
    }

    return entities, nil
}

// GetById retrieves an entity by its ID from the database.
func (r *PostgresBaseRepository[T]) GetById(ctx context.Context, id string) (T, error) {
    query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", r.table)
    row := r.Db.QueryRowContext(ctx, query, id)

    var entity T
    // Scan data into entity, depends on the structure of T
    if err := row.Scan(&entity); err != nil {
        return entity, err
    }

    return entity, nil
}

// Drop deletes the table (use with caution).
func (r *PostgresBaseRepository[T]) Drop(ctx context.Context) error {
    query := fmt.Sprintf("DROP TABLE IF EXISTS %s", r.table)
    _, err := r.Db.ExecContext(ctx, query)
    return err
}