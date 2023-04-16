package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time `json:"startDate" gorm:"type: varchar(255)"`
	EndDate   time.Time `json:"endDate" gorm:"type: varchar(255)"`
	User      User      `json:"user"`
	UserID    int       `json:"userID"`
	Price     int       `json:"price"`
	Status    string    `json:"status" gorm:"type: varchar(255)"`
}
