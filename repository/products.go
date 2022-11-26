package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db}
}

func (p *ProductRepository) AddProduct(product model.Product) error {
	err := p.db.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (p *ProductRepository) ReadProducts() ([]model.Product, error) {
	var products []model.Product

	err := p.db.Where("stock > 0").Find(&products).Error
	if err != nil {
		return []model.Product{}, err
	}

	return products, nil
	// TODO: replace this
}

func (p *ProductRepository) DeleteProduct(id uint) error {
	err := p.db.Delete(&model.Product{}, id).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (p *ProductRepository) UpdateProduct(id uint, product model.Product) error {
	err := p.db.Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}
