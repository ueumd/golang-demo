package models

import (
	orm "myapiserver/api/database"
)

type User struct {
	ID       int64  `json:"id"`       // 列名为 `id`
	Username string `json:"username"` // 列名为 `username`
	Password string `json:"password"` // 列名为 `password`
}

// var Users []User

func (user User) Insert() (id int64, err error)  {

	//添加(创建)数据
	result := orm.Eloquent.Create(&user)

	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func (user *User) Users() (users []User, err error)  {
	if err = orm.Eloquent.Find(&users).Error; err != nil {
		return
	}
	return
}

func (user *User) Update(id int64) (updateUser User, err error)  {
	if err = orm.Eloquent.Select([]string{"id", "username"}).First(&updateUser, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Model(&updateUser).Updates(&user).Error; err != nil {
		return
	}

	return
}

func (user *User) Destroy(id int64) (Result User, err error)  {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}

	Result = *user

	return
}
