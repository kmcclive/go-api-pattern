package goapipattern

type ProductService interface {
	Service[Product]
	SearchByManufacturer(manufacturerID uint) ([]Product, error)
}
