package models

import "fmt"

type Post struct {
	Id          int64  `default:"ID default"`
	Nickname    string `default:"Default Post Object"`
	Description string `default:"Default Description"`
	Date        string `default:"Default Date"`
	Topic       string `default:"Default Topic"`
}

func (p Post) PrintPost() string {
	return fmt.Sprintf("ID: %v\nNickname: %s\nDescription: %s\n", p.Id, p.Nickname, p.Description)
}
