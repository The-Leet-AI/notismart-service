package notification

import (
	"log"
	"notismart-service/internal/db"
	"time"
)

type Notification struct {
	ID      string    `db:"id"`
	UserID  string    `db:"user_id"`
	Content string    `db:"content"`
	SendAt  time.Time `db:"send_at"`
	Status  string    `db:"status"`
}

func ScheduleNotification(userID, content string, sendAt time.Time) error {
	notification := Notification{
		UserID:  userID,
		Content: content,
		SendAt:  sendAt,
		Status:  "Pending",
	}

	_, err := db.DB.NamedExec(`INSERT INTO notifications (user_id, content, send_at, status)
                                VALUES (:user_id, :content, :send_at, :status)`, notification)
	if err != nil {
		log.Printf("Error scheduling notification: %v", err)
		return err
	}
	log.Printf("Notification scheduled for user %s at %v", userID, sendAt)
	return nil
}
