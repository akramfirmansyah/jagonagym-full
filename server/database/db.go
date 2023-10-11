package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/akramfirmansyah/jagona-gym/server/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ConnectDB() {
	var err error

	dns := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Failed connect to database!")
	}

	fmt.Println("Success connect to database")
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate!")
	}

	fmt.Println("Success Migrate Database!")
}

func Seeder() {
	var user models.User
	if err := DB.Where("username = ?", "SuperAdmin").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			pass := "jagonagym"

			hash, err := hashPassword(pass)
			if err != nil {
				panic(err)
			}

			newUser := models.User{
				Username: "SuperAdmin",
				Password: hash,
			}

			DB.Create(&newUser)
		}
	}
}
