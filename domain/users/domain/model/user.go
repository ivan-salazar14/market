package model

import "time"

type User struct {
	UserID      string    `json:"userId,required"`
	Email       string    `json:"email,required"`
	Name        string    `json:"name,required"`
	DateCreated time.Time `json:"-"`
	UserModify  int       `json:"product_user_modify,omitempty"`
	DateModify  time.Time `json:"-"`
}
