package mock

import "tkircsi/restful-template/pkg/models"

type ProductModel struct {
	DB *interface{}
}

func NewProductModel() *ProductModel {
	return &ProductModel{
		DB: nil,
	}
}

var products []*models.Product = []*models.Product{
	{
		ID:   "1",
		Name: "Product 1",
	},
	{
		ID:   "2",
		Name: "Product 2",
	},
}

func (p *ProductModel) Get(id string) (*models.Product, error) {
	var product models.Product
	for _, p := range products {
		if p.ID == id {
			product = *p
			break
		}
	}
	return &product, nil
}

func (p *ProductModel) GetAll() ([]*models.Product, error) {
	return products, nil
}

func (p *ProductModel) Save(prod *models.Product) error {
	products = append(products, prod)
	return nil
}
