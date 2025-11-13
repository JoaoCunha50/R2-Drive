package config

import (
	"api/data/translations"
	"api/data/users"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDBconnection(url string) *gorm.DB {
	newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: 200 * time.Millisecond,
            LogLevel:      logger.Warn,
            Colorful:      true,
        },
    )

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{ Logger: newLogger })
	if err != nil {
		log.Fatal(err)
	}
    
	err = db.Exec(`
		DO $$ 
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
				CREATE TYPE role_enum AS ENUM ('admin', 'member');
			END IF;
		END
		$$;
	`).Error
	if err != nil {
		log.Fatal("Falha ao criar o tipo ENUM 'role_enum':", err)
	}


	db.AutoMigrate(&users.User{}, &translations.Translation{})
	
	return db
}

func CreateAdmin(username string , email string, password string, db *gorm.DB) error {
	var hash, errCrypt = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errCrypt != nil {
		return errCrypt
	}

	var admin = users.User{Username: username, Email: email, Password: string(hash), Name: "João Cunha"}

	err := db.First(&admin, "username = ?", username).Error
	if err == nil {
		return nil
	}

	err1 := db.Create(&admin).Error
	if err1 != nil {
		return err1
	}

	return nil
}