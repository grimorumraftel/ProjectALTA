package users

import (
	"time"

	"gorm.io/gorm"
)

type TopUps struct {
	gorm.Model
	TopUp_ID     uint
	Username     string `gorm:"foreignKey:UserRefer"`
	Amount_TopUp uint
	Topup_At     time.Time
}

type Transfers struct {
	gorm.Model
	Transfer_ID     uint
	Amount_Transfer uint
	Transfer_at     time.Time
}

type Users struct {
	gorm.Model
	Username   string
	Name       string
	Phone      string
	address    string
	password   string
	created_at time.Time
	Balance    uint
}

// FITUR NO.3
func checkAccount(connection *gorm.DB, Username string) ([]Username, error) {
	var Usernames []Users
	err := connection.Table("users_Usernames").Joins("join Usernames on Username = users_Usernames.Username").Where("Usernames = ?", Username).Find(&Usernames).Error
	if err != nil {
		return nil, err
	}

	return Usernames, nil
}

// FITUR NO.4
func updateAccount(connection *gorm.DB, item Item) (bool, error) {
	err := connection.Save(&item).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// FITUR NO.5
func deleteAccount(connection *gorm.DB, Username string) (bool, error) {
	err := connection.Delete(&Users{}, Username).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// FITUR NO.6
func topUpAccount(connection *gorm.DB, Username string, Amount_TopUp uint) (bool, error) {
	usernameTopup := UserItem{Username: Username, Amount_TopUp: Amount_TopUp}
	err := connection.Create(&usernameTopup).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// FITUR NO.7
func transferBalance(connection *gorm.DB, Transfer_ID Transfers) (bool, error) {
	err := connection.Create(&Transfers).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

// FITUR NO.8
func historyTopup()

// FITUR NO.9
func historyTransfer()

// FITUR NO.10
func searchProfile()
