package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

type Item struct {
	gorm.Model
	Name string
}

type UserItem struct {
	gorm.Model
	UserID uint
	ItemID uint
}

func CreateItem(connection *gorm.DB, item Item) (bool, error) {
	err := connection.Create(&item).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetItems(connection *gorm.DB, userID uint) ([]Item, error) {
	var items []Item
	err := connection.Table("user_items").Joins("join items on items.id = user_items.item_id").Where("user_id = ?", userID).Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func UpdateItem(connection *gorm.DB, item Item) (bool, error) {
	err := connection.Save(&item).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteItem(connection *gorm.DB, itemID uint) (bool, error) {
	err := connection.Delete(&Item{}, itemID).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func AddItemToUser(connection *gorm.DB, userID uint, itemID uint) (bool, error) {
	userItem := UserItem{UserID: userID, ItemID: itemID}
	err := connection.Create(&userItem).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
