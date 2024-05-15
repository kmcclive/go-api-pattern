package sql

import (
	"errors"

	"github.com/kmcclive/goapipattern"
	"gorm.io/gorm"
)

type ManufacturerService struct {
	db *gorm.DB
}

func NewManufacturerService(db *gorm.DB) goapipattern.ManufacturerService {
	return &ManufacturerService{
		db: db,
	}
}

func (s *ManufacturerService) Create(manufacturer *goapipattern.Manufacturer) error {
	return s.db.Create(manufacturer).Error
}

func (s *ManufacturerService) Delete(id uint) error {
	var manufacturer goapipattern.Manufacturer

	if resp := s.db.First(&manufacturer, id); errors.Is(resp.Error, gorm.ErrRecordNotFound) {
		return goapipattern.ErrNotFound
	}

	return s.db.Delete(&manufacturer).Error
}

func (s *ManufacturerService) FetchByID(id uint) (*goapipattern.Manufacturer, error) {
	var result goapipattern.Manufacturer

	if resp := s.db.First(&result, id); resp.Error != nil {
		if errors.Is(resp.Error, gorm.ErrRecordNotFound) {
			return nil, goapipattern.ErrNotFound
		}
		return nil, resp.Error
	}

	return &result, nil
}

func (s *ManufacturerService) List() (*[]goapipattern.Manufacturer, error) {
	var result []goapipattern.Manufacturer

	if resp := s.db.Find(&result); resp.Error != nil {
		return nil, resp.Error
	}

	return &result, nil
}

func (s *ManufacturerService) Update(manufacturer *goapipattern.Manufacturer) error {
	var original goapipattern.Manufacturer

	if resp := s.db.First(&original, manufacturer.ID); resp.Error != nil {
		if errors.Is(resp.Error, gorm.ErrRecordNotFound) {
			return goapipattern.ErrNotFound
		}
		return resp.Error
	}

	original.Name = manufacturer.Name

	if resp := s.db.Save(&original); resp.Error != nil {
		return resp.Error
	}

	manufacturer.CreatedAt = original.CreatedAt
	manufacturer.UpdatedAt = original.UpdatedAt

	return nil
}
