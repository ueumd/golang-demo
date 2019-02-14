package model

import (
	orm "myapiserver/api/database"
)

func (user *User) GetUser(username string) (u User, err error)  {
	d := orm.Eloquent.Where("username = ?", username).First(&u)
	return u, d.Error
}

// var Users []User
func (user User) Insert() (id uint64, err error)  {

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

func (user *User) Update(id uint64) (updateUser User, err error)  {
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

func (user *User) Destroy(id uint64) (Result User, err error)  {
	if err = orm.Eloquent.Select([]string{"id"}).First(&user, id).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Delete(&user).Error; err != nil {
		return
	}

	Result = *user

	return
}
