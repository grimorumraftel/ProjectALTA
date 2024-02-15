package users

import (
	"time"

	"gorm.io/gorm"
)

type TopUp struct {
	gorm.Model
	TopUpID   uint
	Username  string `gorm:"foreignKey:UserRefer"`
	Amount    uint
	TopUpTime time.Time
}

type Transfer struct {
	gorm.Model
	TransferID   uint
	Amount       uint
	TransferTime time.Time
}

type User struct {
	gorm.Model
	Username      string
	Name          string
	Jenis_Kelamin string
	Umur          string
	Phone         string
	Address       string
	Password      string
	CreatedAt     time.Time
	Balance       uint
}

func (u *User) ChangePassword(connection *gorm.DB, newPassword string) (bool, error) {
	query := connection.Table("users").Where("username = ?", u.Username).Update("password", newPassword)
	if err := query.Error; err != nil {
		return false, err
	}

	return query.RowsAffected > 0, nil
}

func (u *User) ReadAccount(connection *gorm.DB) (string, error) {
	query := connection.Table("users").Where("username = ?", u.Username).First(u)
	if err := query.Error; err != nil {
		return ("Not found"), err
	}

	return u.Username, nil
}

// func (u *User) UpdateAccount(connection *gorm.DB) (string, error) {
// 	query := connection.Table("users").Where("username = ? AND phone = ?", u.Username, u.Phone).Update(u)
// 	if err := query.Error; err != nil {
// 		return "Phone", err
// 	}

// 	return u.Phone, nil
// }

func (u *User) UpdateAccount(connection *gorm.DB, newPhoneNumber string) (string, error) {
	// Store the old phone number
	oldPhoneNumber := u.Phone

	// Prepare update values with new phone number
	updatePhone := map[string]interface{}{
		"phone": newPhoneNumber,
	}

	// Perform the update query
	query := connection.Table("users").Where("username = ? AND phone = ?", u.Username, oldPhoneNumber).Updates(updatePhone)
	if err := query.Error; err != nil {
		return oldPhoneNumber, err
	}

	// Return the new phone number
	return newPhoneNumber, nil
}

func (u *User) DeleteAccount(connection *gorm.DB, Delete string) (string, error) {
	query := connection.Table("users").Where("Username = ?", u.Username).First(u).Delete(Delete)
	if err := query.Error; err != nil {
		return ("username doesn't exist on database"), err
	}

	return u.Username, nil
}

func (t *TopUp) TopUpAccount(connection *gorm.DB, username string, Amount uint) (uint, error) {
	query1 := connection.Table("top_ups").Create(t)
	if err := query1.Error; err != nil {
		return 0, err
	}
	query2 := connection.Table("users").Where("username = ?", username).Update("balance", Amount)
	if err := query2.Error; err != nil {
		return 0, err
	}

	return t.Amount, nil
}

// func (u *User) TransferBalance(connection *gorm.DB, Transfer uint) (uint error) {
// 	query := connection.Table("users").Where("Username = ?", u.Username)
// 	if err := query.Error; err != nil {
// 		return err
// 	}

// 	return Transfer.TransferID, nil
// }

// func (u *User) HistoryTopUp(connection *gorm.DB, historyTopUp string) (string, error) {
// 	query := connection.Table("users").Where("Username = ?", u.Username).Find(u).Select("History Top Up", historyTopUp)
// 	if err := query.Error; err != nil {
// 		return ("history tidak ditemukan"), err
// 	}

// 	return TopUp.TopUpID, nil
// }

// func (u *User) HistoryTransfer(connection *gorm.DB, historyTransfer string) (bool, error) {
// 	query := connection.Table("users").Where("Username = ?", u.Username).Select("History Transfer", historyTransfer)
// 	if err := query.Error; err != nil {
// 		return false, err
// 	}

// 	return query.RowsAffected > 0, nil
// }

func (u *User) SearchProfile(connection *gorm.DB, Name, Umur, Jenis_Kelamin string) (string, string, string, error) {
	query := connection.Table("users").Where("Username = ?", u.Username).Select("Name", "Umur", "Jenis_Kelamin").Find(u)
	if err := query.Error; err != nil {
		return ("Not Found"), "", "", err
	}

	return u.Name, u.Umur, u.Jenis_Kelamin, nil
}

// func CheckAccount(connection *gorm.DB, username string) (User, error) {
// 	var user User
// 	err := connection.Table("users").Where("username = ?", username).First(&user).Error
// 	if err != nil {
// 		return User{}, err
// 	}

// 	return user, nil
// }

// func UpdateAccount(connection *gorm.DB, user User) (bool, error) {
// 	err := connection.Save(&user).Error
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func DeleteAccount(connection *gorm.DB, username string) (bool, error) {
// 	err := connection.Delete(&User{}, "username = ?", username).Error
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func TopUpAccount(connection *gorm.DB, username string, amount uint) (bool, error) {
// 	topUp := TopUp{
// 		Username:  username,
// 		Amount:    amount,
// 		TopUpTime: time.Now(),
// 	}
// 	err := connection.Create(&topUp).Error
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func TransferBalance(connection *gorm.DB, transfer Transfer) (bool, error) {
// 	err := connection.Create(&transfer).Error
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

// func HistoryTopUp() {
// 	// TODO: Implement historyTopup function
// }

// func HistoryTransfer() {
// 	// TODO: Implement historyTransfer function
// }

// func SearchProfile() {
// TODO: Implement searchProfile function
