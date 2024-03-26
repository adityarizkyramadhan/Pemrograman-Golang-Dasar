package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	// Get product from database
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}
	products := s.database.GetProductData()
	var product entity.Product
	for _, p := range products {
		if p.Name == productName {
			product = p
			break
		}
	}
	if product == (entity.Product{}) {
		return errors.New("product not found")
	}
	cart, err := s.database.GetCartItems()
	if err != nil {
		return err
	}
	cart = append(cart, entity.CartItem{
		ProductName: productName,
		Quantity:    quantity,
		Price:       product.Price,
	})
	return s.database.SaveCartItems(cart)
}

func (s *Service) RemoveCart(productName string) error {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	newCarts := []entity.CartItem{}
	for _, cart := range carts {
		if cart.ProductName == productName {
			continue
		}
		newCarts = append(newCarts, cart)
	}

	if len(newCarts) == len(carts) {
		return errors.New("product not found")
	}

	return s.database.SaveCartItems(newCarts)
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	return s.database.SaveCartItems([]entity.CartItem{})
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	totalPrice := 0
	for _, cart := range carts {
		totalPrice += cart.Price * cart.Quantity
	}

	if money < totalPrice {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	return entity.PaymentInformation{
		TotalPrice:  totalPrice,
		Change:      money - totalPrice,
		MoneyPaid:   money,
		ProductList: carts,
	}, s.ResetCart()
}
