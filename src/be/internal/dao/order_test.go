package dao

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestName(t *testing.T) {
	dsn := "root:1234@tcp(127.0.0.1:13306)/be_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect db: %v", err)
		return
	}
	dao := NewOrderDao(db)
	order := dao.GetById(1)
	log.Println(order)
}
