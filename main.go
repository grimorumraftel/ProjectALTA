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
			fmt.Println("Enter Umur:")
			fmt.Scan(&user.Umur)
			fmt.Println("Enter Jenis Kelamin:")
			fmt.Scan(&user.Jenis_Kelamin)

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
			var newPhone string
			fmt.Println("Masukan no. telp lama")
			fmt.Scanln(&Update.Phone)
			fmt.Println("Masukan no. telp baru")
			fmt.Scanln(&newPhone)

			Phone, err := Update.UpdateAccount(database, newPhone)
			if err != nil && Phone == "Not Found" {
				fmt.Println("Not Found")
			} else {
				fmt.Println(Phone, "Berhasil di ubah")
			}
		case 5:
			var delete users.User
			fmt.Println("Masukan username yang ingin dihapus")
			fmt.Scanln(&delete.Username)
			nama, err := delete.DeleteAccount(database, username)
			if err != nil && nama == "username doesn't exist on database" {
				fmt.Println("username doesn't exist on database")
			} else {
				fmt.Println(user.Username, "Berhasil Dihapus") //POSITIVE RETURN
			}
		case 6:
			//call from users.go
			var topup users.TopUp
			fmt.Println("Masukan jumlah yang ingin ditop-up")
			fmt.Scanln(&topup.Amount)
			fmt.Println("Masukan username yang ingin ditop-up")
			fmt.Scanln(&user.Username)
			jumlah, err := topup.TopUpAccount(database, user.Username, topup.Amount)
			if err != nil && jumlah == 0 {
				return
			}
			fmt.Println(topup.Amount, "Berhasil Ditambah") //POSITIVE RETURN
		case 7:
			// call from method TransferBalance on file users.go
			var transfer users.Transfer
			fmt.Println("Masukan username sender")
			fmt.Scanln(&transfer.SenderUsername)
			fmt.Println("Masukan username receiver")
			fmt.Scanln(&transfer.ReceiverUsername)
			fmt.Println("Masukan jumlah yang ingin ditransfer")
			fmt.Scanln(&transfer.Amount)

			senderUsername, transferredAmount, receiverUsername, err := transfer.TransferBalance(database, transfer.SenderUsername, transfer.Amount, transfer.ReceiverUsername)
			if err != nil && transferredAmount == 0 {
				fmt.Println("Gagal melakukan transfer:")
			} else {
				fmt.Println(senderUsername, "Mengirim Saldo Sebesar", transferredAmount, "Ke Username", receiverUsername)
			}
		// case 8:
		// 	//call from history topup
		// 	var htopup users.User
		// 	fmt.Println("Masukan username yang ingin dilihat historynya")
		// 	fmt.Scanln(&htopup.Username)
		// 	history, amount, TopUpTime err := htopup.HistoryTopUp(database, htopup.TopUpID)
		// 	if err != nil && history == 0 {
		// 		fmt.Println("Not Found")
		// 	} else {
		// 		fmt.Println(history.TopupID,  "Ada di database")
		// 	}
		// case 9:
		//call from history transfer
		case 10:
			//call searchprofile
			var sp users.User
			fmt.Println("Masukan username yang ingin dilihat profilenya")
			fmt.Scanln(&sp.Username)
			profSearch, returnName, returnUmur, err := sp.SearchProfile(database, user.Name, user.Umur, user.Jenis_Kelamin)
			if err != nil && profSearch == "Not Found" {
				fmt.Println("Not Found")
			} else {
				fmt.Println(profSearch, returnName, returnUmur)
			}
		}

		fmt.Println("Exited! Thank you")
	}
}
