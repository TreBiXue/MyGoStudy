package join

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Studentid Studentid // 关联字段
	Age       uint8
	Birthday  *time.Time
}

type Studentid struct {
	No   uint8
	Name string gorm:"unique;foreignKey:StudentName"
}

// Create 插入
func Create(db *gorm.DB) {
	user := User{Age: 18, Studentid: Studentid{No: 2, Name: "bb"}}
	db.Create(&user)
}
