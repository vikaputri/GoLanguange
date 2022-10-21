package models

import (
	"middleware/helpers"

	//"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null" json:"full_name" form:"username" valid:"required-Your full name is required"`
	Email        string        `gorm:"not null" json:"email" form:"email" valid:"required-Your email is required,email-Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required-Your password is required,minstringlength(6)-Password has to have minimum length of 6 characters"`
	Age          int           `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,type(int)~Age must be type of integer, range(8|200)~Age must be more than 8"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"socialMedias"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	//_, errCreate := govalidator.ValidateStruct(u)
	//if errCreate != nil {
	//	err = errCreate
	//	return
	//}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {

	//_, errCreate := govalidator.ValidateStruct(u)

	//if errCreate != nil {
	//	err = errCreate
	//	return
	//}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
