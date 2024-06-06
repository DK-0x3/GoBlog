package database

import (
	"GoBlog/internal/database/models"
	"time"
)

var Users []models.User
var Posts []models.Post
var ActiveUser models.User

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

func GetUsers() *[]models.User {
	return &Users
}

func GetPosts() *[]models.Post {
	return &Posts
}

func AddPost(title, text string) {
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

func UpdatePost(id, title, text string) bool {
    for indx, post := range Posts {
		if post.Id == id {
			if title != "" {
				Posts[indx].Title = title
			}
			if text != "" {
				Posts[indx].Text = text
			}
            return true
		}
	}
    return false
}

func DelPost(ID string) {
	for indx, post := range Posts {
		if post.Id == ID {
			Posts = append(Posts[:indx], Posts[indx+1:]...)
		}
	}
}

func GetActiveUser() *models.User {
    return &ActiveUser
}

func PutActiveUser(user models.User) {
    ActiveUser = user
}

func DelActiveUser() {
    ActiveUser = models.User{}
}

func AddCommentPost(text string, postID string) bool {
	for indx, postItem := range Posts {
		if postItem.Id == postID {
			Posts[indx].Comments = append(Posts[indx].Comments, models.Comment{
				User: ActiveUser,
				Text: text,
				DateTime: time.Now(),
			})
			return true
		}
	}
	return false
}
