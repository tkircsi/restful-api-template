package models

type ProductHandler interface {
	Get(string) (*Product, error)
	GetAll() ([]*Product, error)
	Save(*Product) error
}

type Product struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}
