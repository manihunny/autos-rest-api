package autosqlx

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"main/internal/repositories/autorepository"
	"strings"
)

type AutoSqlx struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *AutoSqlx {
	return &AutoSqlx{
		db: db,
	}
}

func (r *AutoSqlx) Create(ctx context.Context, a *autorepository.Auto) error {
	const q = `
		insert into autos (id, brand, model, mileage, number_of_owners) 
			values (:id, :brand, :model, :mileage, :number_of_owners)
	`
	_, err := r.db.NamedExecContext(ctx, q, a)
	return err
}

func (r *AutoSqlx) Get(ctx context.Context, id string) (*autorepository.Auto, error) {
	const q = `
		select * from autos where id = $1
	`
	a := new(autorepository.Auto)
	err := r.db.GetContext(ctx, a, q, id)
	return a, err
}

func (r *AutoSqlx) List(ctx context.Context) ([]*autorepository.Auto, error) {
	const q = `
		select * from autos
	`
	var list []*autorepository.Auto
	err := r.db.SelectContext(ctx, &list, q)
	return list, err
}

func (r *AutoSqlx) Update(ctx context.Context, a *autorepository.Auto, id string) error {
	const q = `
		update autos set id = $1, brand = $2, model = $3, mileage = $4, number_of_owners = $5 
			where id = $6
	`
	_, err := r.db.ExecContext(ctx, q, a.ID, a.Brand, a.Model, a.Mileage, a.NumberOfOwners, id)
	return err
}

func (r *AutoSqlx) PartialUpdate(ctx context.Context, autoData map[string]interface{}, id string) error {
	q := `
		update autos set 
	`
	for key, value := range autoData {
		q += fmt.Sprintf(`%v = '%v',`, key, value)
	}
	q = strings.TrimRight(q, ",") + fmt.Sprintf(` WHERE id = '%v'`, id)
	_, err := r.db.ExecContext(ctx, q)
	return err
}

func (r *AutoSqlx) Delete(ctx context.Context, id string) error {
	const q = `
		delete from autos where id = $1
	`
	_, err := r.db.ExecContext(ctx, q, id)
	return err
}
