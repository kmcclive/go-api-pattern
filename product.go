package goapipattern

type Product struct {
	Model
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	ManufacturerID uint          `json:"manufacturerId"`
	Manufacturer   *Manufacturer `json:"manufacturer"`
}

type ProductService interface {
	Service[Product]
	SearchByManufacturer(manufacturerID uint) ([]Product, error)
}
