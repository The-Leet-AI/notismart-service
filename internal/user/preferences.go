package user

import (
	"log"
	"notismart-service/internal/db"
)

type UserPreference struct {
	UserID          string `db:"user_id"`
	PreferredTime   string `db:"preferred_time"`
	PreferredMethod string `db:"preferred_method"`
}

func GetUserPreferences(userID string) (*UserPreference, error) {
	var prefs UserPreference
	err := db.DB.Get(&prefs, "SELECT * FROM user_preferences WHERE user_id = $1", userID)
	if err != nil {
		log.Printf("Error fetching preferences for user %s: %v", userID, err)
		return nil, err
	}
	return &prefs, nil
}
