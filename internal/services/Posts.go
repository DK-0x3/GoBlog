package services

import (
	"GoBlog/internal/database"
	"GoBlog/internal/database/models"

	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

var iUser database.LocalUser
var iPost database.LocalPost
var iActiveUser database.LocalActiveUser
var iComments database.LocalComment

type ManagementPost interface {
	ViewMyPosts()
	ViewAllPosts()
	AddPosts()
}

func ViewMyPosts() {
	var fx bool = true
	var myPost []models.Post
	posts := *iPost.GetPosts()
	activUser := iActiveUser.GetActiveUser()

	for _, post := range posts {
			if post.Author.Name == activUser.Name {
				myPost = append(myPost, post)
			}
		}
	var j int = len(myPost) - 1
	PrintColorText(color.FgHiYellow, "\nВаши посты:\n\n")

	for fx {
		myPost = nil
		for _, post := range posts {
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
		PrintColorText(color.FgHiCyan, "   " + myPost[j].DateTime.Format("02.01.2006 15:04")+"\n")
		count := 0
		for _, s := range myPost[j].Text {
			if count > 100 {
				if string(s) == " " {
					count = 0
					fmt.Print("\n")
					continue
				}
			}
			fmt.Print(string(s))
			count++
		}

		fmt.Printf("\n\nКомментарии: %d ", len(myPost[j].Comments))
		for _, komment := range myPost[j].Comments {
			PrintColorText(color.FgGreen, "\n  " + komment.User.Name)
			fmt.Print("\n    " + komment.Text + "\n    " + komment.DateTime.Format("02.01.2006 15:04"))
		}

		fmt.Print("\n\n[0] - Вернуться в профиль\n[1] - Прошлый пост\n[2] - Следующий пост\n[3] - Изменить пост\n[4] - Удалить пост\n\n")
		
		var input string
		fmt.Scan(&input)
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
			if newTitle == "0" {
				continue
			}
			fmt.Print("\nИзменение Текста, введи новый заголовок\n*Если не нужно изменять введи '0'\n ")
			newText = ReadInput()
			if newText == "0" {
				continue
			}

			err := iPost.UpdatePost(myPost[j].Id, newTitle, newText)
			if err != nil {
				PrintColorText(color.FgRed, "Ошибка добавления поста\n")
				time.Sleep(2 * time.Second)
				continue
			}
			
			PrintColorText(color.FgGreen, "Пост успешно изменен!\n")
			time.Sleep(2 * time.Second)

		}else if input == "4" {
			
			fmt.Print("\nПост: " + myPost[j].Title + "\nУдалить этот пост? (д/н): ")
			var del string
			fmt.Scan(&del)
			if del == "д" {
				iPost.DelPost(myPost[j].Id)
				fmt.Print("Пост успешно удален!\n")
				time.Sleep(2 * time.Second)
			}else {
				fmt.Print("Удаление отменено\n")
				time.Sleep(2 * time.Second)
			}
		}else {
			fx = false
			break
		}

	}
}

func ViewAllPosts() {
	var fx bool = true
	var post []models.Post = *iPost.GetPosts()

	var j int = len(post) - 1

	PrintColorText(color.FgHiYellow, "\nНедавние посты:\n\n")

	for fx {
		PrintColorText(color.FgHiMagenta, "\n" + strings.ToUpper(post[j].Title))
		PrintColorText(color.FgHiCyan, "   " + post[j].DateTime.Format("02.01.2006 15:04") + "\n")
		count := 0
		for _, s := range post[j].Text {
			if count > 100 {
				count = 0
				fmt.Print("\n")
			}
			fmt.Print(string(s))
			count++
		}
		PrintColorText(color.FgHiBlue, "\nАвтор: "+post[j].Author.Name)

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
		
		var input string
		fmt.Scan(&input)

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
				
				fmt.Print("\nВведи комментарий: \n")
				inputText := ReadInput()
				iComments.AddCommentPost(inputText, post[j].Id)
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

		iPost.AddPost(titlePost, textPost)

		fmt.Print("Пост успешно создан!\n[0] - Вернуться в профиль\n[1] - Просмотр моих постов\n")

		var inp string
		fmt.Scan(&inp)
		if inp == "1" {
			ViewMyPosts()
			break
		} else {
			break
		}
				
	}
}
