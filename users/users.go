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
	Username  string
	Name      string
	Phone     string
	Address   string
	Password  string
	CreatedAt time.Time
	Balance   uint
}

func CheckAccount(connection *gorm.DB, username string) (User, error) {
	var user User
	err := connection.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateAccount(connection *gorm.DB, user User) (bool, error) {
	err := connection.Save(&user).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteAccount(connection *gorm.DB, username string) (bool, error) {
	err := connection.Delete(&User{}, "username = ?", username).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func TopUpAccount(connection *gorm.DB, username string, amount uint) (bool, error) {
	topUp := TopUp{
		Username:  username,
		Amount:    amount,
		TopUpTime: time.Now(),
	}
	err := connection.Create(&topUp).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func TransferBalance(connection *gorm.DB, transfer Transfer) (bool, error) {
	err := connection.Create(&transfer).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func HistoryTopUp() {
	// TODO: Implement historyTopup function
}

func HistoryTransfer() {
	// TODO: Implement historyTransfer function
}

func SearchProfile() {
	// TODO: Implement searchProfile function
}
