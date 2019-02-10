package database

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql" //_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
	"fmt"
)

var Eloquent *gorm.DB

func init()  {
	var err error
	Eloquent, err = gorm.Open("mysql","root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local&timeout=10ms")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}