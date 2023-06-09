package main

import (
	"errors"
	"fmt"
)

type Service interface {
	GetAll() ([]Product, error)
	Store(name, category, token string, count int, price float64) (Product, error)
	Create(data Product, token string) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
	GetByID(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetByID(id int) (Product, error) {
	product, err := s.repository.GetByID(id)
	if err != nil {
		return Product{}, fmt.Errorf("produto não encontrado: %v", err)
	}
	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func validarToken(token string) bool {
	validToken := "1234"
	return token == validToken
}

// cria um novo produto
func (s *service) Create(data Product, token string) (Product, error) {

	if !validarToken(token) {
		return Product{}, errors.New("Token inválido")
	}

	if data.Name == "" {
		return Product{}, errors.New("O campo 'name' é obrigatório")
	}
	if data.Category == "" {
		return Product{}, errors.New("O campo 'category' é obrigatório")
	}
	if data.Count <= 0 {
		return Product{}, errors.New("O campo 'count' deve ser maior que zero")
	}
	if data.Price <= 0 {
		return Product{}, errors.New("O campo 'price' deve ser maior que zero")
	}

	data.ID = productIDCounter
	productIDCounter++

	newProduct, err := s.repository.Create(data)
	if err != nil {
		return Product{}, err
	}

	return newProduct, nil
}

// func store armazena um novo produto
func (s *service) Store(name, category, token string, count int, price float64) (Product, error) {
	// Verificar se o token é válido
	if !validarToken(token) {
		return Product{}, fmt.Errorf("você não tem permissão para fazer a solicitação solicitada")
	}

	// Criar um novo produto com ID gerado automaticamente
	newProduct := Product{
		ID:       generateID(),
		Name:     name,
		Category: category,
		Count:    count,
		Price:    price,
	}

	// Armazenar o novo produto no repositório
	product, err := s.repository.Store(newProduct.Name, newProduct.Category, newProduct.Count, newProduct.Price)
	if err != nil {
		return Product{}, fmt.Errorf("erro ao armazenar o produto: %v", err)
	}

	return product, nil
}

// Função Update para atualizar um produto pelo ID
func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	// Buscar o produto pelo ID no repositório
	product, err := s.repository.GetByID(id)
	if err != nil {
		return Product{}, fmt.Errorf("produto não encontrado: %v", err)
	}

	// Atualizar os campos do produto com os novos valores
	product.Name = name
	product.Category = productType
	product.Count = count
	product.Price = price

	// Atualizar o produto no repositório
	updatedProduct, err := s.repository.Update(product.ID, product.Name, product.Category, product.Count, product.Price)
	if err != nil {
		return Product{}, fmt.Errorf("erro ao atualizar o produto: %v", err)
	}

	return updatedProduct, nil
}

// Função UpdateName para atualizar o nome de um produto pelo ID
func (s *service) UpdateName(id int, name string) (Product, error) {
	// Buscar o produto pelo ID no repositório
	product, err := s.repository.GetByID(id)
	if err != nil {
		return Product{}, fmt.Errorf("produto não encontrado: %v", err)
	}

	// Atualizar o nome do produto
	product.Name = name

	// Atualizar o produto no repositório
	updatedProduct, err := s.repository.Update(product.ID, product.Name, product.Category, product.Count, product.Price)
	if err != nil {
		return Product{}, fmt.Errorf("erro ao atualizar o produto: %v", err)
	}

	return updatedProduct, nil
}

// Função Delete para excluir um produto pelo ID
func (s *service) Delete(id int) error {
	// Buscar o produto pelo ID no repositório
	_, err := s.repository.GetByID(id)
	if err != nil {
		return fmt.Errorf("produto não encontrado: %v", err)
	}

	// Excluir o produto do repositório
	err = s.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("erro ao excluir o produto: %v", err)
	}

	return nil
}
