package v1

import domain "presentation/internal/domain/application"

type requestConverter struct {
}

var RequestConverter = requestConverter{}

func (requestConverter) ApplicationCreate(r applicationReq) domain.Application {
	return domain.Application{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Age:       r.Age,
		Phone:     r.Phone,
		Car: domain.Car{
			Brand:      r.Car.Brand,
			Model:      r.Car.Model,
			Power:      r.Car.Power,
			IsElectric: r.Car.IsElectric,
		},
	}
}
