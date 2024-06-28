package autorepository

import "context"

type Auto struct {
	ID             string  `json:"id" db:"id"`
	Brand          string  `json:"brand" db:"brand"`
	Model          string  `json:"model" db:"model"`
	Mileage        float64 `json:"mileage" db:"mileage"`
	NumberOfOwners int     `json:"number_of_owners" db:"number_of_owners"`
}

type AutoRepository interface {
	Create(ctx context.Context, a *Auto) error
	Get(ctx context.Context, id string) (*Auto, error)
	List(ctx context.Context) ([]*Auto, error)
	Update(ctx context.Context, a *Auto, id string) error
	PartialUpdate(ctx context.Context, a map[string]interface{}, id string) error
	Delete(ctx context.Context, id string) error
}
