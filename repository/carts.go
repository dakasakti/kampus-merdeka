package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return CartRepository{db}
}

func (c *CartRepository) ReadCart() ([]model.JoinCart, error) {
	var joins []model.JoinCart

	err := c.db.Raw("SELECT c.id, c.product_id, p.name, sum(c.quantity) as quantity, sum(c.total_price) as total_price FROM carts as c JOIN products AS p ON p.id = c.product_id WHERE C.deleted_at IS NULL GROUP BY c.id, c.product_id, p.name, quantity, total_price").Scan(&joins).Error
	if err != nil {
		return []model.JoinCart{}, err
	}

	return joins, nil
	// TODO: replace this
}

func (c *CartRepository) AddCart(product model.Product) error {
	product.Stock -= 1
	totalPrice := product.Price - ((product.Discount * 0.01) * product.Price)

	var cart model.Cart
	err := c.db.Where("product_id = ?", product.ID).Find(&cart).Error
	if err != nil {
		return err
	}

	// update cart
	cart.ProductID = product.ID
	cart.Quantity += 1
	cart.TotalPrice += totalPrice

	err = c.db.Save(&product).Error
	if err != nil {
		return err
	}

	err = c.db.Save(&cart).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (c *CartRepository) DeleteCart(id uint, productID uint) error {
	// product
	var product model.Product
	err := c.db.Where("id = ?", productID).Find(&product).Error
	if err != nil {
		return err
	}

	product.Stock += 1
	err = c.db.Save(&product).Error
	if err != nil {
		return err
	}

	// cart
	var cart model.Cart
	err = c.db.Where("id = ? AND product_id = ?", id, productID).Find(&cart).Error
	if err != nil {
		return err
	}

	if cart.Quantity > 1 {
		cart.TotalPrice -= (cart.TotalPrice / cart.Quantity)
		cart.Quantity -= 1

		err = c.db.Save(&cart).Error
		if err != nil {
			return err
		}

		return nil
	}

	err = c.db.Delete(&cart).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}

func (c *CartRepository) UpdateCart(id uint, cart model.Cart) error {
	err := c.db.Where("id = ?", id).Updates(&cart).Error
	if err != nil {
		return err
	}

	return nil
	// TODO: replace this
}
