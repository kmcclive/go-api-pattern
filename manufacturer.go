package goapipattern

type Manufacturer struct {
	Model
	Name string `json:"name" binding:"required"`
}

type ManufacturerService interface {
	Service[Manufacturer]
}
