package entity

import "time"

type UserPermission struct {
	ID     int
	Method string
	Qty    int
	DateTo time.Time
	UserID int
}
