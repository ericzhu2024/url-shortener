package models

import "time"

type Link struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Code      string     `json:"code" gorm:"uniqueIndex;size:10"`
	Original  string     `json:"original" gorm:"index"`
	Clicks    int        `json:"clicks" gorm:"default:0"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"` // pointer allows null
}