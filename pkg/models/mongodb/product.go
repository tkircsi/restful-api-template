package mongodb

import (
	"context"
	"log"
	"time"
	"tkircsi/restful-template/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductModel struct {
	DB *mongo.Database
}

func NewProductModel(host string) (*ProductModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(host))
	if err != nil {
		return nil, err
	}
	return &ProductModel{DB: client.Database("Product")}, nil
}

func (p *ProductModel) GetAll() ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cursor, err := p.DB.Collection("Products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var products []*models.Product
	if err = cursor.All(ctx, &products); err != nil {
		log.Println(err)
		return nil, err
	}
	return products, nil
}

func (p *ProductModel) Get(id string) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var product *models.Product
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := p.DB.Collection("Products").FindOne(ctx, bson.M{"_id": oid}).Decode(&product); err != nil {
		log.Println(err)
		return nil, err
	}
	return product, nil
}

func (p *ProductModel) Save(prod *models.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := p.DB.Collection("Products").InsertOne(ctx, prod)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res)
	return nil
}
