package model

import (
	"loginimpl/dao"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"username"`
	Password string `form:"password"`
}

func CreateAUser(user *User) (err error) {
	err = dao.Db.Create(&user).Error
	return
}

func GetAUser(username string) (user *User, err error) {
	user = new(User)

	if err = dao.Db.Debug().Where(" name = ?", username).Find(&user).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAUserPassword(user *User) (err error) {
	err = dao.Db.Debug().Save(&user).Error
	return
}

func DeleteAUser(username string) (err error) {
	err = dao.Db.Debug().Where("name = ?", username).Delete(User{}).Error
	return
}
