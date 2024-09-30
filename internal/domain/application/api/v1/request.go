package v1

type applicationReq struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name" binding:"min=2,max=255,inputNameFormat"`
	LastName  string `json:"last_name" binding:"min=3,max=255,inputNameFormat"`
	Phone     string `json:"phone"`
	Age       int    `json:"age"`
	Car       car    `json:"car"`
}

type car struct {
	Brand      string `json:"brand" binding:"oneof=BMW HYUNDAI TOYOTA"`
	Model      string `json:"model"`
	Power      string `json:"power"`
	IsElectric bool   `json:"is_electric"`
}
