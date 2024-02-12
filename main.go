package main

import (
	"be21/config"
	"be21/users"
	"fmt"

	"gorm.io/gorm"
)

func printMenu() {
	fmt.Println("Pilih menu")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("99. Exit")
}

func registerUser(connection *gorm.DB, user users.User) (bool, error) {
	err := connection.Create(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func loginUser(connection *gorm.DB, username string, password string) (bool, error) {
	var user users.User
	err := connection.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // User not found
		}
		return false, err // Database error
	}

	return true, nil // User found
}

func main() {
	database := config.InitMysql()
	config.Migrate(database)
	var input int

	var username, password string
	var user users.User
	for input != 99 {
		printMenu()
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
		switch input {
		case 1:
			fmt.Print("Enter username:")
			fmt.Scanln(&username)
			fmt.Print("Enter password:")
			fmt.Scanln(&password)
			success, err := loginUser(database, username, password)
			if err != nil {
				fmt.Println("Error logging in:", err)
			} else if !success {
				fmt.Println("Invalid username or password.")
			} else {
				fmt.Println("Logged in successfully.")
			}
		case 2:
			fmt.Print("Enter username:")
			fmt.Scanln(&user.Username)
			fmt.Print("Enter password:")
			fmt.Scanln(&user.Password)
			success, err := registerUser(database, user)
			if err != nil {
				fmt.Println("Error registering:", err)
			} else if !success {
				fmt.Println("Registration failed.")
			} else {
				fmt.Println("Registered successfully.")
			}
		}
	}
	fmt.Println("Exited! Thank you")
}
