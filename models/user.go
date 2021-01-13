package models

import "time"

type User struct {
	Id          int64     `orm:"auto"`
	Name        string    `orm:"null;size(20)"`
	UserName    string    `orm:"size(15)"`
	Birthday    time.Time `orm:"size(100)"`
	PhoneNumber string    `orm:"size(40)"`
	Password    string    `orm:"size(100)"`
	JwtToken    string    `orm:"null;size(350)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}
