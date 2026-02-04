package db

import "time"

type User struct {
	ID        int64     `json:"id"         db:"id"`
	TGUserID  int64     `json:"tg_id"      db:"tg_id"`      // ID из Telegram
	Name      string    `json:"name"       db:"name"`       // Отображаемое имя: "Иван Математика"
	Role      string    `json:"role"       db:"role"`       // "teacher", "student", "parent"
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Время создания пользователя
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
