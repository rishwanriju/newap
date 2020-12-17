package entity

type News struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Post  string `json:"post" binding:"max=100" gorm:"type:varchar(100)"`
}
