package goapipattern

type Product struct {
	Model
	Name           string        `json:"name" binding:"required"`
	Description    string        `json:"description"`
	ManufacturerID uint          `json:"manufacturerId" binding:"required"`
	Manufacturer   *Manufacturer `json:"manufacturer"`
}

type ProductService interface {
	Service[Product]
	SearchByManufacturer(manufacturerID uint) (*[]Product, error)
}
