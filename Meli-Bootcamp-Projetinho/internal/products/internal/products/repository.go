package products

import (
	"fmt"

	"github.com/vraissa/Meli-bootcamp-projetinho/pkg/store"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Count    int     `json:"count"`
	Price    float64 `json:"price"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(name, category string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	err := r.db.Read(&ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}
	ultimoProduto := ps[len(ps)-1]
	return ultimoProduto.ID, nil
}

func (r *repository) Store(id int, name, productType string, count int, price float64) (Product, error) {
	produtos := []Product{}
	r.db.Read(&produtos)
	p := Product{id, name, productType, count, price}
	produtos = append(produtos, p)
	if err := r.db.Write(produtos); err != nil {
		return Product{}, err
	}
	return p, nil

}

func (repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Category: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil

}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d não encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	return nil
}
