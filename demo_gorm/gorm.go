package demo_gorm
/**
https://github.com/hatlonely/microservices/blob/master/internal/comment_like/comment_like.go
 */

import (
	"time"

	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"strings"
	"github.com/spaolacci/murmur3"
	"fmt"
)



type View struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64
	CreatedAt time.Time
}


var db *gorm.DB

func testDB1()  {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//创建表
	if !db.HasTable(&View{}) {
		if err :=
		// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句  // 为模型`User`创建表
			db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&View{}).Error; err != nil{
			panic(err)
		}
	}
}

func testDB2(username, password, addr, name string) {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")
	fmt.Println(config)
}

func init()  {
	testDB1()

}

func main()  {
	// testDB2("root", "123456", "127.0.0.1", "test")
}


func DoView(ip, ua, title string) error  {
	view :=&View{
		Ip: 		ip,
		Ua:			ua,
		Title: 		title,
		Hash:		murmur3.Sum64([]byte(strings.Join([]string{ip, ua, title}, "-"))) >> 1,
		CreatedAt:	time.Now(),
	}

	if err := db.Create(view).Error; err != nil {
		return  err
	}

	return  nil
}


func Add(a, b int) int {
	return a + b
}
