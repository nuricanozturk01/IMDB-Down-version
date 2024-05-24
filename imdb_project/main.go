package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"imdb_project/entity"
)

func main() {
	//app.Run()

	dsn := "root:12345678@tcp(127.0.0.1:3306)/imdb_clone?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Callback().Create().Before("gorm:create").Register("before_create", func(db *gorm.DB) {
		if db.Statement.Schema != nil {
			ctx := context.Background()
			for _, field := range db.Statement.Schema.Fields {
				if field.Name == "ID" && field.DataType == "uuid" {
					value, _ := field.ValueOf(ctx, db.Statement.ReflectValue)
					if value == nil || value == uuid.Nil {
						err := field.Set(ctx, db.Statement.ReflectValue, uuid.New())
						if err != nil {
							return
						}
					}
				}
			}
		}
	})

	// Veritabanı migrasyonları
	db.AutoMigrate(&entity.User{}, &entity.Movie{}, &entity.TVShow{}, &entity.Watchlist{}, &entity.WatchListItem{}, &entity.Rating{}, &entity.Photo{}, &entity.Trailer{}, &entity.Celebrity{}, &entity.Company{})

	// CRUD işlemleri
	user := createUser(db)
	readUser(db)
	updateUser(db, user.ID)
	//deleteUser(db, updatedUser.ID)
}

func createUser(db *gorm.DB) entity.User {
	user := entity.User{
		ID:        uuid.New(),
		Username:  "nuri",
		Password:  "password123",
		FirstName: "John",
		LastName:  "Doe",
		Email:     "nuri@example.com",
	}
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
	} else {
		fmt.Println("User created successfully:", user.ID)
	}
	return user
}

func readUser(db *gorm.DB) {
	var user entity.User
	result := db.First(&user)
	if result.Error != nil {
		fmt.Println("Error reading user:", result.Error)
	} else {
		fmt.Println("User read successfully:", user)
	}
}

func updateUser(db *gorm.DB, userID uuid.UUID) entity.User {
	var user entity.User
	result := db.First(&user, "id = ?", userID)
	if result.Error != nil {
		fmt.Println("Error finding user:", result.Error)
		return user
	}

	// Güncelleme işlemi
	result = db.Model(&user).Where("id = ?", userID).Updates(entity.User{LastName: "UpdatedLastName"})
	if result.Error != nil {
		fmt.Println("Error updating user:", result.Error)
	} else {
		fmt.Println("User updated successfully:", user)
	}
	return user
}

func deleteUser(db *gorm.DB, userID uuid.UUID) {
	var user entity.User
	result := db.First(&user, "id = ?", userID)
	if result.Error != nil {
		fmt.Println("Error finding user:", result.Error)
		return
	}

	result = db.Delete(&user, "id = ?", userID.String())
	if result.Error != nil {
		fmt.Println("Error deleting user:", result.Error)
	} else {
		fmt.Println("User deleted successfully")
	}
}
