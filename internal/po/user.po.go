package po

import "time"

type User struct {
	ID       int      `gorm:"primaryKey, autoIncrement, not null; unique;"`
	Email    string   `gorm:"size:50; not null; unique"`
	Password string   `gorm:"not null"`
	UserInfo UserInfo `gorm:"foreignKey:UserID;references:ID"`
}

type UserInfo struct {
	ID          int       `gorm:"primaryKey, autoIncrement, not null, unique;"`
	UserID      int       `gorm:"not null"`
	FullName    string    `gorm:"size:255; not null"`
	Gender      bool      `gorm:"not null"`
	BirthDay    time.Time `gorm:"not null"`
	Address     string    `gorm:"not null"`
	Phone       string    `gorm:"size:10, not null"`
	Description string    `gorm:"not null"`
}

func (u *User) TableName() string {
	return "user"
}

func (ui *UserInfo) TableName() string {
	return "user_info"
}
