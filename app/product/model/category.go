package model

import (
	"context"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `gorm:"many2many:t_product_category;"` // many to many
}

func (Category) TableName() string {
	return "t_category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) CategoryQuery {
	return CategoryQuery{ctx: ctx, db: db}
}

func (c CategoryQuery) GetProductsByCategoryName(name string) (categorys []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).
		Where("name = ?", name).
		Preload("Products").
		Find(&categorys).Error
	return
}
