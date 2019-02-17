package main

import "github.com/jinzhu/gorm"

type Category struct {
	Name        string `gorm:"primary_key"`
	Description string `gorm:"size:255;default:'nothing in here'"`
}

type Email struct {
	ID         int
	UserId     int
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

type Origin struct {
	ID        int
	ProductID uint
	Address1  string `gorm:"not null;unique"`
	Address2  string `gorm:"unique"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"`
	Code string `gorm:"index:idx_name_code"`
}

type Product struct {
	gorm.Model
	Name string `gorm:"index;size:255"`

	Sid         int        `gorm:"unique_index"`
	Categories  []Category `gorm:"many2many:categories_product;"`
	Emails      []Email    `gorm:"ForeignKey:UserId"`
	Origin      *Origin
	Languages   []Language
	Score       *float64 `gorm:"not null;default:1.0"`
	Description string   `gorm:"size:255;default:'nothing in here'"`
}

type GreekAlphabet struct {
	ID         uint   `gorm:"primary_key"`
	LatinName  string `gorm:"unique_index"`
	UpperCode  rune
	LowerCode  rune
	IsFrequent bool `gorm:"index"`
}

func main()  {
	db, err := gorm.Open("mysql", "root:123456@/myapiserver?charset=utf8&parseTime=True&loc=Local")

	// defer db.close()

	categories := []Category{
		Category{"mobile phone", "a hand-held mobile radiotelephone for use in an area divided into small sections, each with its own short-range transmitter/receiver"},
		Category{"apple", ""},
	}
	emails := []Email{Email{Email: "example@domain.com", Subscribed: false}}
	origin := Origin{Address1: "apple company address", Address2: "test"}
	languages := []Language{Language{Name: "中国", Code: "cn"}, Language{Name: "美国", Code: "us"}}
	score := float64(0.0)
	product := Product{
		Name:       "iphone7",
		Sid:        1211,
		Categories: categories,
		Emails:     emails,
		Origin:     &origin,
		Languages:  languages,
		Score:      &score,
	}

	err = db.Create(&product).Error
	if err != nil {
		// t.Error(err)
	}
}
