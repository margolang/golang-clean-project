package response

import domain "presentation/internal/domain/application"

type applicationConverter struct{}

var ApplicationConverter = applicationConverter{}

type application struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

type car struct {
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Power      string `json:"power"`
	IsElectric bool   `json:"is_electric"`
}

func (c *applicationConverter) Response(r domain.Application) application {
	return application{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Phone:     r.Phone,
	}
}
