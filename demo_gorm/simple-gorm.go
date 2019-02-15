package main

import (
		_"github.com/jinzhu/gorm/dialects/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"fmt"
)

type StudyUser struct {
	ID       int64  // 列名为 `id`
	Username string // 列名为 `username`
	Password string // 列名为 `password`
}

func (StudyUser) TableName() string  {
	return  "studyusers"
}

var db *gorm.DB
var err error

func initDB()  {
	db, err = gorm.Open("mysql", "root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	defer db.Close()

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
}

func testDB1()  {
	db, err = gorm.Open("mysql", "root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	defer db.Close()

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	// 建表
	db.CreateTable(&StudyUser{})

	// 添加数据
	user := StudyUser{Username:"root", Password: "123456"}
	result := db.Create(&user)

	if result.Error != nil {
		fmt.Printf("insert row err %v", result.Error)
		return
	}

	fmt.Println(user.ID) //返回id


	//查询单条数据
	getUser := StudyUser{}
	db.Select([]string{"id", "username"}).First(&getUser, 2)
	fmt.Println(getUser) //打印查询数据

	//修改数据
	//user.Username = "update username"
	//user.Password = "update password"
	//db.Save(&user)

	//查询列表数据
	users := []StudyUser{}
	db.Find(&users)
	fmt.Println(&users) //获取所有数据

	//删除数据
	//db.Delete(&user)
}

func testDB2()  {
	db, err = gorm.Open("mysql", "root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	// 将实际执行的sql打印出来 true | false
	db.LogMode(true)
	// db.Debug().Select([]string{"id", "username", "password"}).First(&user) 打印单行

	defer db.Close()

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	user := StudyUser{}
	// 获取第一条记录，按主键排序
	//  SELECT * FROM studyusers ORDER BY id LIMIT 1;
	db.Debug().Select([]string{"id", "username", "password"}).First(&user)
	fmt.Println(user) //{2 update username update password}

	user2 := StudyUser{}
	// 使用主键获取记录
	// SELECT * FROM studyusers where id = 4
	db.Debug().Select([]string{"id", "username", "password"}).First(&user2, 4)
	fmt.Println(user2) //{4 root 123456}

	// 获取所有记录
	users := []StudyUser{}
	db.Debug().Find(&users)
	fmt.Println(users) //[{2 update username update password} {3 root 123456} {4 root 123456} {5 root 123456} {6 root 123456} {7 root 123456}]

	//https://segmentfault.com/a/1190000003036452
	//db.Exec("select * from studyusers limit 10")

}

func main()  {
	//testDB1()
	testDB2()
}