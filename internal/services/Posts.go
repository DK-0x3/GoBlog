package services

import (
	"GoBlog/internal/database"
	"GoBlog/internal/database/models"

	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func ViewMyPosts() {
	var fx bool = true
	var myPost []models.Post
	posts := database.GetPosts()
	activUser := database.GetActiveUser()

	for _, post := range *posts {
			if post.Author.Name == activUser.Name {
				myPost = append(myPost, post)
			}
		}
	var j int = len(myPost) - 1
	PrintColorText(color.FgHiYellow, "\nВаши посты:\n\n")

	for fx {
		myPost = nil
		for _, post := range *posts {
			if post.Author.Name == activUser.Name {
				myPost = append(myPost, post)
			}
		}
		if len(myPost) <= 0 {
			fmt.Print("У вас нет постов!")
			time.Sleep(3 * time.Second)
			break
		}
		if j < 0 || j > len(myPost)-1 {
			j = len(myPost) - 1
		}
		PrintColorText(color.FgHiMagenta, strings.ToUpper(myPost[j].Title))
		PrintColorText(color.FgHiCyan, "\n"+myPost[j].DateTime.Format("02.01.2006 15:04"))
		fmt.Print("\n\n" + myPost[j].Text)
		fmt.Printf("\n\nКомментарии: %d ", len(myPost[j].Comments))
		for _, komment := range myPost[j].Comments {
			PrintColorText(color.FgGreen, "\n  " + komment.User.Name)
			fmt.Print("\n    " + komment.Text + "\n    " + komment.DateTime.Format("02.01.2006 15:04"))
		}

		fmt.Print("\n\n[0] - Вернуться в профиль\n[1] - Прошлый пост\n[2] - Следующий пост\n[3] - Изменить пост\n[4] - Удалить пост\n\n")
		
		input := ReadInput()
		if input == "1" {
			if j <= 0 {
				j = len(myPost) - 1
			} else {
				j--
			}

		} else if input == "2" {
			if j >= len(myPost)-1 {
				j = 0
			} else {
				j++
			}
		}else if input == "3" {
			var newTitle string
			var newText string

			fmt.Print("\nИзменение Заголовка, введи новый заголовок\n*Если не нужно изменять введи '0'\n ")
			newTitle = ReadInput()

			fmt.Print("\nИзменение Текста, введи новый заголовок\n*Если не нужно изменять введи '0'\n ")
			newText = ReadInput()

			WritePost(models.Post{
				Id: myPost[j].Id,
				Title: newTitle,
				Text: newText,
			})

			PrintColorText(color.FgGreen, "Пост успешно изменен!\n")
			time.Sleep(2 * time.Second)

		}else if input == "4" {
			var del string
			fmt.Print("\nПост: " + myPost[j].Title + "\nУдалить этот пост? (д/н): ")
			del = ReadInput()
			if del == "д" {
				DelPost(myPost[j].Id)
				
				fmt.Print("Пост успешно удален!\n")
			}else {
				fmt.Print("Удаление отменено\n")
			}

		}else {
			fx = false
			break
		}

	}
}

func ViewAllPosts() {
	var fx bool = true
	var post []models.Post = *database.GetPosts()

	var j int = len(post) - 1

	PrintColorText(color.FgHiYellow, "\nНедавние посты:\n\n")

	for fx {
		PrintColorText(color.FgHiMagenta, "\n" + strings.ToUpper(post[j].Title))
		PrintColorText(color.FgHiCyan, "\n" + post[j].DateTime.Format("02.01.2006 15:04"))
		fmt.Print("\n\n" + post[j].Text)
		fmt.Printf("\n\nКомментарии: %d ", len(post[j].Comments))
		for _, komment := range post[j].Comments {
			PrintColorText(color.FgGreen, "\n  " + komment.User.Name)
			fmt.Print("\n    " + komment.Text + "\n    " + komment.DateTime.Format("02.01.2006 15:04"))
		}
		if database.ActiveUser.Name == "" {
			fmt.Print("\n\n[0] - Вернуться\n[1] - Прошлый пост\n[2] - Следующий пост\n")
		}else {
			fmt.Print("\n\n[0] - Вернуться\n[1] - Прошлый пост\n[2] - Следующий пост\n[3] - Написать комментарий\n")
		}
		
		input := ReadInput()

		if input == "1" {
			if j <= 0 {
				j = len(post) - 1
			} else {
				j--
			}
		} else if input == "2" {
			if j >= len(post)-1 {
				j = 0
			} else {
				j++
			}
		}else if input == "3" {
			if database.ActiveUser.Name != "" {
				var inputText string
				fmt.Print("\nВведи комментарий: ")
				inputText = ReadInput()
				AddCommentPost(inputText, post[j].Id)
				fmt.Print("Ваш комментарий успешно добавлен!")
				time.Sleep(2 * time.Second)
			}
		}else {
			fx = false
			break
		}
	}
}

func AddPosts() {
	for {
		var titlePost string
		var textPost string
		posts := database.GetPosts()

		PrintColorText(color.FgGreen, "Создание Нового поста\n*[0] - Вернуться в профиль\n\n")
		
		fmt.Print("Введи Заголовок поста: ")
		titlePost = ReadInput()
		if titlePost == "0" {
			break
		}

		fmt.Print("Введи Текст поста: ")
		textPost = ReadInput()
		if textPost == "0" {
			break
		}

		*posts = append(*posts, models.Post{
			Id:       database.ActiveUser.Name + "_" + string(len(*database.GetPosts())),
			Title:    titlePost,
			Text:     textPost,
			DateTime: time.Now(),
			Author:   database.ActiveUser,
		})

		fmt.Print("Пост успешно создан!\n[0] - Вернуться в профиль\n[1] - Просмотр моих постов\n")

		inp := ReadInput()
		if inp == "1" {
			ViewMyPosts()
			break
		} else {
			break
		}
				
	}
}

func WritePost(postInfo models.Post) {
	posts := database.GetPosts()
	postsWrite := *posts
	for indx, post := range *posts {
		if post.Id == postInfo.Id {
			if postInfo.Title != "0" {
				postsWrite[indx].Title = postInfo.Title
			}
			if postInfo.Text != "0" {
				postsWrite[indx].Text = postInfo.Text
			}
		}
	}
}

func DelPost(postID string) {
	posts := database.GetPosts()
	for indx, post := range *posts {
		if post.Id == postID {
			*posts = append((*posts)[:indx], (*posts)[indx+1:]...)
		}
	}
}