package models

type User struct {
	Id   int64 `orm:"auto"`
	Name string
}
