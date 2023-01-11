package usermodel

type User struct {
	Id       int    `json:"id,omitempty" gorm:"column:id"`
	Login    string `json:"login" gorm:"column:login"`
	Password string `json:"password" gorm:"column:password"`
	Lang     string `json:"lang" gorm:"column:lang"`
}

func (User) TableName() string {
	return "res_users"
}

type UserUpdate struct {
	Login    *string `json:"login" gorm:"column:login"`
	Password *string `json:"password" gorm:"column:password"`
	Lang     *string `json:"lang" gorm:"column:lang"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserCreate struct {
	Login    string `json:"login" gorm:"column:login"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}
