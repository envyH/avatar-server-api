package service

import (
	"avatar/internal/db"
	"context"
)

func GetNotification(key string) (string, error) {
	rows, err := db.GetDB().Query(context.Background(), "SELECT value FROM notifications WHERE key = $1", key)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var value string
	if rows.Next() {
		if err := rows.Scan(&value); err != nil {
			return "", err
		}
		return value, nil
	}

	return "", nil
}
