package convert

import (
	domain "presentation/internal/domain/application"
	"presentation/internal/model"
)

type application struct {
}

var Application = application{}

func (application) Model(d domain.Application) model.Application {
	return model.Application{
		ID:         d.ID,
		FirstName:  d.FirstName,
		LastName:   d.LastName,
		Age:        d.Age,
		CarBrand:   d.Car.Brand,
		CarPower:   d.Car.Power,
		CarModel:   d.Car.Model,
		IsElectric: d.Car.IsElectric,
	}
}

func (a *application) DomainSlice(m model.ApplicationSlice) []domain.Application {
	res := make([]domain.Application, len(m))
	for _, v := range m {
		res = append(res, a.Domain(v))
	}

	return res
}

func (application) Domain(m *model.Application) domain.Application {
	return domain.Application{
		ID:        m.ID,
		Phone:     m.Phone,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Car: domain.Car{
			Brand:      m.CarBrand,
			Power:      m.CarPower,
			Model:      m.CarModel,
			IsElectric: m.IsElectric,
		},
	}
}
