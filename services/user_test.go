package services

import (
	"log"
	"testing"

	"github.com/Siddheshk02/portfolio-management/models"
	"github.com/Siddheshk02/portfolio-management/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupTestDB() (*gorm.DB, error) {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test-db password=Sid@2002 sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}

func TestRegisterUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	userRepo := repositories.NewUserRepository(db)
	userService := NewUserService(userRepo)

	err = userService.RegisterUser("Test User", "password")
	assert.NoError(t, err)

	user, err := userRepo.GetUserByUsername("Test User")
	assert.NoError(t, err)
	assert.Equal(t, "Test User", user.Username)

	defer db.DropTable("users")
}

func TestAuthenticateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	db.AutoMigrate(&models.User{})

	userRepo := repositories.NewUserRepository(db)
	userService := NewUserService(userRepo)

	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	userRepo.CreateUser(&models.User{Username: "Test User", Password: string(hashedPassword)})

	token, err := userService.AuthenticateUser("Test User", password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	defer db.DropTable("users")
}
