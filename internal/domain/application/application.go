package application

import "context"

type Application struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	Age       int
	Car       Car
}

type Car struct {
	Brand      string
	Model      string
	Power      string
	IsElectric bool
}

type ApplicationService interface {
	Create(ctx context.Context, app Application) error
	Update(ctx context.Context, app Application) error
	Get(ctx context.Context, id string) (Application, error)
	GetAll(ctx context.Context) ([]Application, error)
	Delete(ctx context.Context, id string) error
}

type ApplicationRepository interface {
	Add(ctx context.Context, app Application) error
	Update(ctx context.Context, app Application) error
	Get(ctx context.Context, id string) (Application, error)
	GetAll(ctx context.Context) ([]Application, error)
	Delete(ctx context.Context, id string) error
}
