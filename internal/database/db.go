package database

import (
	"GoBlog/internal/database/models"
	"errors"
	"time"
)

type StartApp interface {
	GetUsers()
	GetPosts()
	AddPost(title, text string)
	UpdatePost(id, title, text string)
	DelPost(ID string)
	GetActiveUser()
	PutActiveUser(user models.User)
	DelActiveUser()
	AddCommentPost(text string, postID string)
}

var Users []models.User
var Posts []models.Post
var ActiveUser models.User

type LocalUser struct {
	User models.User
}
type LocalPost struct {
	Post models.Post
}
type LocalActiveUser struct {
	ActiveUser models.User
}
type LocalComment struct {
	Comment models.Comment
}

func InitDB() {
    comment := []models.Comment{}
    comment = append(comment, models.Comment{
                User: models.User{
		            Name: "xxxx",
		            Email: "@.x",
		            Password: "xxx",
	            },
                Text: "[pdgfpkpof]",
                DateTime: time.Now(),
            })

    Posts = append(Posts, models.Post{
        Id: "xxxx_1",
        Title: "xz",
        Text: "хз зачем это надо",
        DateTime: time.Now(),
        Author: models.User{Name: "xxxx", Email: "zzz@com"},
        Comments: comment,
    })
    Posts = append(Posts, models.Post{
        Id: "xxxx_2",
        Title: "ccc",
        Text: "sdfsdghfdvо",
        DateTime: time.Now(),
        Author: models.User{Name: "xxxx", Email: "zzz@com"},
    })
	Users = append(Users, models.User{
		Name: "xxxx",
		Email: "@.x",
		Password: "xxx",
	})
}

func (u *LocalUser) GetUsers() *[]models.User {
	return &Users
}

func (p *LocalPost) GetPosts() *[]models.Post {
	return &Posts
}

func (p *LocalPost) AddPost(title, text string) {
    if title == "" {
        title = "Пустой загаловок"
    }
    if text == "" {
        text = "Нет контента"
    }
    Posts = append(Posts, models.Post{
		Id:       ActiveUser.Name + "_" + string(len(Posts)),
		Title:    title,
		Text:     text,
		DateTime: time.Now(),
		Author:   ActiveUser,
	})
}

func (p *LocalPost) UpdatePost(id, title, text string) error {
    for indx, post := range Posts {
		if post.Id == id {
			if title != "" {
				Posts[indx].Title = title
			}
			if text != "" {
				Posts[indx].Text = text
			}
            return nil
		}
	}
    return errors.New("ошибка: пост не найден")
}

func (p *LocalPost) DelPost(ID string) {
	for indx, post := range Posts {
		if post.Id == ID {
			Posts = append(Posts[:indx], Posts[indx+1:]...)
		}
	}
}

func (p *LocalActiveUser) GetActiveUser() *models.User {
    return &ActiveUser
}

func (p *LocalActiveUser) PutActiveUser(user models.User) {
    ActiveUser = user
}

func (p *LocalActiveUser) DelActiveUser() {
    ActiveUser = models.User{}
}

func (p *LocalComment) AddCommentPost(text string, postID string) error {
	for indx, postItem := range Posts {
		if postItem.Id == postID {
			Posts[indx].Comments = append(Posts[indx].Comments, models.Comment{
				User: ActiveUser,
				Text: text,
				DateTime: time.Now(),
			})
			return nil
		}
	}
	return errors.New("ошибка добавления комментария")
}
