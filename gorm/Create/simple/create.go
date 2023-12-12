package simple

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name     string
	Age      uint8
	Birthday *time.Time
}

// Create 插入
func Create(db *gorm.DB) {
	/* 全插入*/
	currentTime := time.Now()
	users := []*User{
		&User{Name: "Jinzhu2", Age: 18, Birthday: &currentTime},
		&User{Name: "Jackson1", Age: 19, Birthday: nil},
	}

	result := db.Create(users) // 传递切片以插入多行数据

	// result.RowsAffected  返回插入记录的条数
	fmt.Printf("User Insrt End; Count: %v, Error: %v\n", result.RowsAffected, result.Error)

	/* 指定字段插入插入*/
	user2 := User{
		Name:     "user2",
		Age:      18,
		Birthday: &currentTime,
	}

	result = db.Select("Name", "Age", "Birthday").Create(&user2)
	// result.RowsAffected  返回插入记录的条数
	fmt.Printf("(Select)User Insrt End; Count: %v, Error: %v\n", result.RowsAffected, result.Error)

	/* 忽略字段插入插入*/
	user3 := User{
		Name:     "user3",
		Age:      18,
		Birthday: &currentTime,
	}
	result = db.Omit("Birthday").Create(&user3)

	// result.RowsAffected  返回插入记录的条数
	fmt.Printf("(Omit)User Insrt End; Count: %v, Error: %v\n", result.RowsAffected, result.Error)

	/* 批量插入:①批次插入性能高（db交互次数少）   ②比起单条插入是一个事务，某条失败可回滚 */
	var userList []User

	for i := 1500; i >= 1001; i-- {
		name := fmt.Sprintf("xiaoming%d", i)
		user1 := User{
			Name:     name,
			Age:      18,
			Birthday: nil,
		}
		userList = append(userList, user1)
	}

	result = db.CreateInBatches(userList, 100)
	// result.RowsAffected  返回插入记录的条数
	fmt.Printf("User CreateInBatches End; Count: %v, Error: %v\n", result.RowsAffected, result.Error)
}

// BeforeCreate 钩子
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// 在创建记录之前执行的操作
	currentTime := time.Now()
	tenYearsAgo := currentTime.Add(-10 * 365 * 24 * time.Hour)
	u.Birthday = &tenYearsAgo
	fmt.Printf("Hook Create!")

	return nil
}
