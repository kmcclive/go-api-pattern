package goapipattern

type Manufacturer struct {
	Model
	Name string `json:"name"`
}

type ManufacturerService interface {
	Service[Manufacturer]
}
