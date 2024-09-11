package models

type User struct {
	Id       int64  `default:"ID default"`
	Nickname string `default:"Default Post Object"`
	Email    string `default:"Default Description"`
	Date     string `default:"Default Date"`
}
