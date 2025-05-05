package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Question struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	Title       string
	Description string
	Lang        string
	Difficulty  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Lang struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Answer struct {
	ID        uint `gorm:"primaryKey"`
	QID       uint
	Answer    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
