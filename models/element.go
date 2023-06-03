package models

import "time"

type Element struct {
	Id          uint       `gorm:"primaryKey"`
	WaterStatus string     `gorm:"not null;type:varchar(191)"`
	WindStatus  string     `gorm:"not null;type:varchar(191)"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
