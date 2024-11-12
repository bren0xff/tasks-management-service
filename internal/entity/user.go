package entity

type User struct {
	ID       string `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
