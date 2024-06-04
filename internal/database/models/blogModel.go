package models

import "time"

type Post struct {
	Id 		 string
	Title    string
	Text     string
	DateTime time.Time
	Author User
	Comments []Comment
}

type Comment struct {
	User User
	Text string
	DateTime time.Time
}