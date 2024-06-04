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

func GetActiveUser() *models.User {
    return &ActiveUser
}

func PutActiveUser(user models.User) {
    ActiveUser = user
}

func DelActiveUser() {
    ActiveUser = models.User{}
}
