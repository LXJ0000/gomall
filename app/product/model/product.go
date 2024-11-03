package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categorys []Category `gorm:"many2many:t_product_category;"` // many to many
}

func (Product) TableName() string {
	return "t_product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB) ProductQuery {
	return ProductQuery{ctx: ctx, db: db}
}

func (p ProductQuery) GetByID(ID int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, ID).Error
	return
}

func (p ProductQuery) SearchProduct(q string) (products []Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).
		Where("name LIKE ? or desc LIKE ?", "%"+q+"%", "%"+q+"%").
		Find(&products).Error
	return
}
