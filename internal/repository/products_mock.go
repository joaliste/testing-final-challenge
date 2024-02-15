package repository

import "app/internal"

func NewProductsMock() *ProductsMock {
	return &ProductsMock{}
}

type ProductsMock struct {
	SearchProductsFunc func(query internal.ProductQuery) (p map[int]internal.Product, err error)

	Spy struct {
		SearchProducts int
	}
}

func (pm *ProductsMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	pm.Spy.SearchProducts++
	return pm.SearchProductsFunc(query)
}
