package services

import (
	"GoBlog/internal/database"
	"GoBlog/internal/database/models"
	"time"
)

func AddCommentPost(text string, postID string) bool {
	posts := database.GetPosts()
	user := database.GetActiveUser()
	for indx, postItem := range *posts {
		if postItem.Id == postID {
			post := *posts
			post[indx].Comments = append(post[indx].Comments, models.Comment{
				User: *user,
				Text: text,
				DateTime: time.Now(),
			})
			return true
		}
	}
	return false
}