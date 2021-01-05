package models

import "time"

type User struct {
	Id          int64     `orm:"auto"`
	Name        string    `orm:"size(20)"`
	PhoneNumber int       `orm:"size(40)"`
	Password    string    `orm:"size(100)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}
