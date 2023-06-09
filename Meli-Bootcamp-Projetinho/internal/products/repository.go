package products

import (
	"errors"
	"fmt"
)

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
	Create(data Product) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
	GetByID(id int) (Product, error)
}

type repository struct {
	products []Product
}

// NewRepository cria uma nova instância do repositório
func NewRepository() Repository {
	return &repository{
		products: []Product{},
	}
}

// GetAll retorna todos os produtos armazenados no repositório
func (r *repository) GetAll() ([]Product, error) {
	return r.products, nil
}

// Store armazena um novo produto no repositório
func (r *repository) Store(name, category string, count int, price float64) (Product, error) {
	product := Product{
		ID:       generateID(),
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}
	r.products = append(r.products, product)
	return product, nil
}

// Create cria um novo produto com os dados fornecidos e o armazena no repositório
func (r *repository) Create(data Product) (Product, error) {
	data.ID = generateID()
	r.products = append(r.products, data)
	return data, nil
}

// Update atualiza os dados de um produto com o ID fornecido
func (r *repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	for i, p := range r.products {
		if p.ID == id {
			r.products[i].Name = name
			r.products[i].Category = productType
			r.products[i].Count = count
			r.products[i].Price = price
			return r.products[i], nil
		}
	}
	return Product{}, errors.New("Produto não encontrado")
}

// UpdateName atualiza o nome de um produto com o ID fornecido
func (r *repository) UpdateName(id int, name string) (Product, error) {
	for i, p := range r.products {
		if p.ID == id {
			r.products[i].Name = name
			return r.products[i], nil
		}
	}
	return Product{}, errors.New("Produto não encontrado")
}

// Delete remove um produto com o ID fornecido do repositório
func (r *repository) Delete(id int) error {
	for i, p := range r.products {
		if p.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return errors.New("Produto não encontrado")
}

var productIDCounter int

func generateID() int {
	productIDCounter++
	return productIDCounter
}

func (r *repository) GetByID(id int) (Product, error) {
	for _, p := range r.products {
		if p.ID == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("produto não encontrado")
}
