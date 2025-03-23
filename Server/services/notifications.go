package services

import (
	"github.com/roughtomato/notifyMe/models"
	"gorm.io/gorm"
)

func GetAllNotificationsByUserId(db *gorm.DB, userId int) ([]models.Notification, error) {
	var notifications []models.Notification
	err := db.Where("user_id = ? AND status != ?", userId, "read").Find(&notifications).Error
	return notifications, err
}

func AddNotification(db *gorm.DB, notification models.Notification) error {
	return db.Create(&notification).Error
}
