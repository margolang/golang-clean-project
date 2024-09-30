package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	errors2 "presentation/internal/domain/application/errors"
	"presentation/internal/model"

	domain "presentation/internal/domain/application"
	"presentation/internal/domain/application/db/postgres/convert"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type applicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) domain.ApplicationRepository {
	return &applicationRepository{db: db}
}

func (a applicationRepository) Add(ctx context.Context, app domain.Application) error {
	m := convert.Application.Model(app)

	err := m.Insert(ctx, a.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (a applicationRepository) Update(ctx context.Context, app domain.Application) error {
	// TODO
	return nil
}

func (a applicationRepository) Get(ctx context.Context, id string) (domain.Application, error) {
	res, err := model.Applications(model.ApplicationWhere.ID.EQ(id)).One(ctx, a.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Application{}, errors2.ErrApplicationNotFound
		}
		return domain.Application{}, err
	}

	return convert.Application.Domain(res), nil
}

func (a applicationRepository) GetAll(ctx context.Context) ([]domain.Application, error) {
	// TODO
	return nil, nil
}

func (a applicationRepository) Delete(ctx context.Context, id string) error {
	nj, err := model.Applications(model.ApplicationWhere.ID.EQ(id)).DeleteAll(ctx, a.db)
	if err != nil {
		fmt.Print(nj)
	}
	return err
}
