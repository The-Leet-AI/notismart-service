package notification

import (
	"log"
	"notismart-service/internal/db"
	"notismart-service/internal/user"
	"time"
)

func DispatchPendingNotifications() {
	var pendingNotifications []Notification

	// Fetch notifications where the send time has passed and status is still 'Pending'
	err := db.DB.Select(&pendingNotifications, "SELECT * FROM notifications WHERE send_at <= $1 AND status = 'Pending'", time.Now())
	if err != nil {
		log.Printf("Error fetching pending notifications: %v", err)
		return
	}

	// Dispatch each notification
	for _, notification := range pendingNotifications {
		prefs, err := user.GetUserPreferences(notification.UserID)
		if err != nil {
			log.Printf("Error fetching user preferences for user %s: %v", notification.UserID, err)
			continue
		}

		if prefs.PreferredMethod == "Email" {
			// Implement SendEmailNotification
			// SendEmailNotification(notification.UserID, notification.Content)
		} else if prefs.PreferredMethod == "SMS" {
			// Implement SendSMSNotification
		} else if prefs.PreferredMethod == "Push" {
			// Implement SendPushNotification
		}

		// Mark the notification as 'Sent'
		db.DB.Exec("UPDATE notifications SET status = 'Sent' WHERE id = $1", notification.ID)
		log.Printf("Notification %s sent to user %s", notification.ID, notification.UserID)
	}
}
