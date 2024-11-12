package entity

type User struct {
	ID       string `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password string `json:"-" gorm:"type:varchar(255)"`
	Role     Role   `json:"role" gorm:"type:varchar(50)"`
}
