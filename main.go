package main

import (
	"be21/config"
	"be21/users"
	"fmt"

	"gorm.io/gorm"
)

func printMenu() {
	fmt.Println("Pilih menu")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Read Account")
	fmt.Println("4. Update Account")
	fmt.Println("5. Delete Account")
	fmt.Println("6. Top-up")
	fmt.Println("7. Transfer")
	fmt.Println("8. History Top-up")
	fmt.Println("9. History Transfer")
	fmt.Println("10. Search Profile")
	fmt.Println("0. Exit")
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
			fmt.Println("Enter Username:")
			fmt.Scan(&user.Username)
			fmt.Println("Enter Password:")
			fmt.Scan(&user.Password)
			fmt.Println("Enter Name:")
			fmt.Scan(&user.Name)
			fmt.Println("Enter Phone:")
			fmt.Scan(&user.Phone)
			fmt.Println("Enter Address:")
			fmt.Scan(&user.Address)

			success, err := registerUser(database, user)
			if err != nil {
				fmt.Println("Error registering:", err)
			} else if !success {
				fmt.Println("Registration failed.")
			} else {
				fmt.Println("Registered successfully.")
			}
		case 2:
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
		case 3:
			//call function ReadAccount from users.go
			var usr users.User
			fmt.Println("Masukan username yang ingin dilihat/dibaca")
			fmt.Scanln(&usr.Username)
			nama, err := usr.ReadAccount(database)
			if err != nil && nama == "Not Found" {
				fmt.Println("Not Found")
			} else {
				fmt.Println(nama, "Ada di database")
			}

		case 4:
			//call from users.go
			var Update users.User
			fmt.Println("Masukan no. telp lama")
			fmt.Scanln(&Update.Phone)
			Phone, err := Update.UpdateAccount(database)
			if err != nil && Phone == "Not Found" {
				fmt.Println("Not Found")
			} else {
				fmt.Println(Phone, "Berhasil di ubah")
			}
		case 5:
			//call from users.go
		case 6:
			//call from users.go
		case 7:
			//call from users.go
		case 8:
			//call from users.go
		case 9:
			//call from users.go
		case 10:
			//call from users.go
		}
	}
	fmt.Println("Exited! Thank you")
}
