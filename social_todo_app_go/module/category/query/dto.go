package query

type CategoryDTO struct {
	Id   string `gorm:"column:id;" json:"id"`
	Name string `gorm:"column:name;" json:"name"`
}

func (CategoryDTO) TableName() string {
	return "category"
}
