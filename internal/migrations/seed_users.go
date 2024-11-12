package migrations

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"tasksManagement/internal/entity"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []entity.User{
		{ID: uuid.NewString(), Name: "Carlos Manager", Email: "carlos@example.com", Password: "manager123", Role: "manager"},
		{ID: uuid.NewString(), Name: "Joao Manager", Email: "joao@example.com", Password: "manager123", Role: "manager"},
		{ID: uuid.NewString(), Name: "John Technician", Email: "john.tech@example.com", Password: "technician123", Role: "technician"},
		{ID: uuid.NewString(), Name: "Jane Technician", Email: "jane.tech@example.com", Password: "technician123", Role: "technician"},
	}

	for _, user := range users {
		user.Password = HashPassword(user.Password)
		if err := db.Create(&user).Error; err != nil {
			log.Printf("Failed to insert user %s: %v", user.Email, err)
			return err
		}
		log.Printf("User %s inserted successfully", user.Email)
	}
	return nil
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
