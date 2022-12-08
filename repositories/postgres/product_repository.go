package postgres

import (
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/entity"
)

type ProductRepository struct {
}

func (s *ProductRepository) Store(p entity.Produto) (entity.Produto, error) {
	db := database.GetDatabase()
	err := db.Create(&p).Error

	if err != nil {
		return entity.Produto{}, err
	}

	return p, nil
}

func (s *ProductRepository) Delete(id uint) (entity.Produto, error) {
	db := database.GetDatabase()
	n, err := s.Get(id)
	if err != nil {
		return entity.Produto{}, err
	}

	err = db.Delete(&n).Error
	if err != nil {
		return entity.Produto{}, fmt.Errorf("cannot delete file: %v", err)
	}

	return n, nil
}

func (s *ProductRepository) Update(p entity.Produto) error {
	db := database.GetDatabase()

	err := db.Save(&p).Error
	if err != nil {
		return fmt.Errorf("cannot update client on pg: %v", err)
	}

	return nil
}

func (s *ProductRepository) Get(id uint) (entity.Produto, error) {
	db := database.GetDatabase()
	var p entity.Produto
	err := db.First(&p, id).Error

	if err != nil {
		return entity.Produto{}, fmt.Errorf("cannot find product by id: %v", err)
	}

	return p, nil
}

func (s *ProductRepository) GetAll() ([]entity.Produto, error) {
	db := database.GetDatabase()
	var p []entity.Produto
	err := db.Find(&p).Error

	if err != nil {
		return nil, fmt.Errorf("cannot find products: %v", err)
	}

	return p, err
}
