package sql

import (
	"errors"

	"github.com/kmcclive/goapipattern"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductService struct {
	db                  *gorm.DB
	manufacturerService goapipattern.ManufacturerService
}

func NewProductService(db *gorm.DB, manufacturerService goapipattern.ManufacturerService) goapipattern.ProductService {
	return &ProductService{
		db:                  db,
		manufacturerService: manufacturerService,
	}
}

func (s *ProductService) Create(product *goapipattern.Product) error {
	manufacturer, err := s.manufacturerService.FetchByID(product.ManufacturerID)
	if err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			return goapipattern.ErrParentNotFound
		}
		return err
	}

	if resp := s.db.Create(product); resp.Error != nil {
		return resp.Error
	}

	product.Manufacturer = manufacturer

	return nil
}

func (s *ProductService) Delete(id uint) error {
	var product goapipattern.Product

	if resp := s.db.First(&product, id); errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return goapipattern.ErrNotFound
	}

	return s.db.Delete(&product).Error
}

func (s *ProductService) FetchByID(id uint) (*goapipattern.Product, error) {
	var result goapipattern.Product

	if resp := s.db.Preload(clause.Associations).First(&result, id); errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return nil, goapipattern.ErrNotFound
	}

	return &result, nil
}

func (s *ProductService) List() (*[]goapipattern.Product, error) {
	return s.list(nil)
}

func (s *ProductService) ListByManufacturer(manufacturerID uint) (*[]goapipattern.Product, error) {
	criteria := goapipattern.Product{
		ManufacturerID: manufacturerID,
	}

	return s.list(&criteria)
}

func (s *ProductService) list(criteria *goapipattern.Product) (*[]goapipattern.Product, error) {
	db := s.db.Preload(clause.Associations)

	if criteria != nil {
		db = db.Where(criteria)
	}

	var result []goapipattern.Product

	if resp := db.Find(&result); resp.Error != nil {
		return nil, resp.Error
	}

	return &result, nil
}

func (s *ProductService) Update(product *goapipattern.Product) error {
	var original goapipattern.Product

	if resp := s.db.First(&original, product.ID); errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return goapipattern.ErrNotFound
	}

	manufacturer, err := s.manufacturerService.FetchByID(product.ManufacturerID)
	if err != nil {
		if errors.Is(err, goapipattern.ErrNotFound) {
			return goapipattern.ErrParentNotFound
		}
		return err
	}

	original.Name = product.Name
	original.Description = product.Description
	original.ManufacturerID = product.ManufacturerID

	if resp := s.db.Save(&original); resp.Error != nil {
		return resp.Error
	}

	product.CreatedAt = original.CreatedAt
	product.UpdatedAt = original.UpdatedAt
	product.Manufacturer = manufacturer

	return nil
}
