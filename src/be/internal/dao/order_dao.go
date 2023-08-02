package dao

import (
	"gorm.io/gorm"

	"be/internal/model"
)

type OrderDao struct {
	db *gorm.DB
}

func NewOrderDao(db *gorm.DB) *OrderDao {
	return &OrderDao{db: db}
}

func (d *OrderDao) Create(o *model.Order) {
	d.db.Create(o)
}

func (d *OrderDao) GetById(id int64) *model.Order {
	o := &model.Order{Id: id}
	d.db.First(o)
	return o
}
