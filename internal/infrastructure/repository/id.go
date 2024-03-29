package repository

import (
	"context"
	utlits "github.com/3110Y/cc-utlits"
	"github.com/3110Y/idsaver/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type IdRepository struct {
	db *sqlx.DB
}

func NewIdRepository(db *sqlx.DB) *IdRepository {
	return &IdRepository{db: db}
}

func (i IdRepository) Add(ctx context.Context, id entity.IDEntity) (*uint64, error) {
	dsn := "INSERT INTO id (id, uid, len, allow) VALUES (:id, :uid, :len, :allow)"
	result, err := i.db.NamedExecContext(ctx, dsn, &id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	return utlits.Pointer(uint64(rowsAffected)), err
}

func (i IdRepository) Item(ctx context.Context, id string) (*entity.IDEntity, error) {
	pr := entity.IDEntity{}
	err := i.db.GetContext(ctx, &pr, "SELECT * FROM id WHERE id = $1", id)
	return &pr, err
}
