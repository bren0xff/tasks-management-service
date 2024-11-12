package entity

type Task struct {
	ID      string `json:"id" gorm:"type:char(36);primaryKey"`
	Summary string `json:"summary" gorm:"type:varchar(255)"`
	Date    string `json:"date" gorm:"type:date"`
	UserID  string `json:"user_id" gorm:"type:char(36);not null"`
}
