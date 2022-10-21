package models

//"github.com/asaskevich/govalidator"

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url" valid:"required~Photo url is required"`
	UserID   uint   `json:user_id`
	User     *User
}

//func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
//	_, errCreate := govalidator.ValidateStruct(p)

//	if errCreate != nil {
//		err = errCreate
//		return
//	}

//	err = nil
//	return
//}

//func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
//	_, errCreate := govalidator.ValidateStruct(p)

//	if errCreate != nil {
//		err = errCreate
//		return
//	}

//	err = nil
//	return
//}
