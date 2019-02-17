package database

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/spf13/viper"
)

var Eloquent *gorm.DB

func Init() *gorm.DB  {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.dbname"),
	)
}

func openDB(username, password, addr, dbname string) *gorm.DB  {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		dbname,
		true,
		"Local",
		)

	var err error
	//Eloquent, err = gorm.Open(
	//	"mysql",
	//	"root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	Eloquent, err = gorm.Open(
		"mysql",
		config)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}

	return Eloquent
}