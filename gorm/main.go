package main

import (
	"MyGoStudy/gorm/Create/join"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mygostudy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Printf("DB Connect Failed!")
	}

	//CreateSimple.Create(db)
	join.Create(db)
}
