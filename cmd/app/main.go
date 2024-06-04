package main

import (
	"GoBlog/internal/handler"
	"GoBlog/internal/models"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
    var Users []models.User
    var Posts []models.Post

    Posts = append(Posts, models.Post{
        Title: "xz",
        Text: "хз зачем эьто надо",
        DateTime: time.Now(),
        Author: models.User{Name: "xxxx", Email: "zzz@com"},
    })
    Posts = append(Posts, models.Post{
        Title: "ccc",
        Text: "sdfsdghfdvо",
        DateTime: time.Now(),
        Author: models.User{Name: "xxxx", Email: "zzz@com"},
    })

    for {
        clearScreen()
        Title := "Приветствую тебя\n\n[1] - Войти\n[2] - Зарегистрироваться\n[0] - завершить сеанс\nВведи команду: "
        PrintColorText(color.FgBlue, Title)
        
        var action string
        fmt.Scan(&action)

        switch action {
        case "0":
            clearScreen()
            fmt.Println("Завершение программы...")
            os.Exit(0)
        case "1":
            for {
                clearScreen()
                var email string
                var password string
                PrintColorText(color.FgGreen, "Вход\n\n*Вернуться назад - [0]\n\n")
                
                fmt.Print("Введи Email: ")
                fmt.Scan(&email)
                if (email == "0") {
                    break
                }

                fmt.Print("Введи Пароль: ")
                fmt.Scan(&password)
                if (password == "0") {
                    break
                }

                activUser, err := handler.Entrance(email, password, &Users)
                if err != "" {
                    var input string
                    PrintColorText(color.FgRed, "Ошибка: " + err)
                    fmt.Print("\n[1] - Попробовать снова\n[0] - Вернуться на главную\nВаш выбор: ")
                    fmt.Scan(&input)
                    if input == "0" {
                        break
                    }else if (input == "1") {
                        continue
                    }else {
                        break
                    }
                }
                for {
                    clearScreen()
                    PrintColorText(color.FgGreen, "Добро пожаловать " + activUser.Name)
                    fmt.Print("\n[0] - Выход из профиля\n[1] - Просмотр и редактирование постов\n[2] - Создать новый пост\n")
                    var inputProfile string
                    fmt.Scan(&inputProfile)
                    if inputProfile == "0" {
                        break
                    }else if (inputProfile == "1") {

                    }else if (inputProfile == "2") {
                        clearScreen()
                        for {
                            var titlePost string
                            var textPost string

                            PrintColorText(color.FgGreen, "Создание Нового поста\n*[0] - Вернуться в профиль\n\n")
                            fmt.Print("Введи Заголовок поста: ")
                            fmt.Scan(&titlePost)
                            if titlePost == "0" {
                                break
                            }
                            fmt.Print("Введи Текст поста: ")
                            fmt.Scan(&textPost)
                            if textPost == "0" {
                                break
                            }

                            Posts = append(Posts, models.Post{
                                Title: titlePost,
                                Text: textPost,
                                DateTime: time.Now(),
                                Author: activUser,
                                Like: []models.Like{},
                            })
                            fmt.Print("Пост успешно создан!\n[0] - Вернуться в профиль\n[1] - Просмотр моих постов\n")
                            var inp string
                            fmt.Scan(&inp)
                            if inp == "1" {
                                ViewMyPosts(activUser, &Posts)
                                break
                            }else {
                                break
                            }
                        }
                    }
                }
                
            }
        case "2":
            for {
                clearScreen()
                var email string
                var name string
                PrintColorText(color.FgGreen, "Регистрация\n\n*Вернуться назад - [0]\n\n")
                
                fmt.Print("Введи Email: ")
                fmt.Scan(&email)
                if (email == "0") {
                    break
                }

                fmt.Print("Введи Name: ")
                fmt.Scan(&name)
                if (name == "0") {
                    break
                }
                result, registrValid := handler.Registration(email, name, &Users)
                if registrValid {

                    PrintColorText(color.FgGreen, "Успешная регистрация!\nОсталось войти в профиль\n\n")
                    fmt.Printf("Ваш пароль: %s\n\nВведи что угодно для перехода на главную...\n", result)
                    fmt.Scan(&result)
                    break

                }else {

                    var input string
                    PrintColorText(color.FgRed, "Ошибка регистрации: " + result)
                    fmt.Print("\n[1] - Попробовать снова\n[0] - Вернуться на главную\nВаш выбор: ")
                    fmt.Scan(&input)
                    if input == "0" {
                        break
                    }else if (input == "1") {
                        continue
                    }else {
                        break
                    }
                }

            }
        default:
            fmt.Print("Неверный ввод, такой команды не найдено")
        }
    }
    
	
}

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}

func PrintColorText(colorText color.Attribute, text string) {
    color.Set(colorText)

    fmt.Print(text)

    color.Unset()
}

func ViewMyPosts(activUser models.User, posts *[]models.Post) {
    var fx bool = true
    var myPost []models.Post

    for _, post := range *posts {
        if (post.Author.Name == activUser.Name) {
            myPost = append(myPost, post)
        }
    }
    var j int = len(myPost) - 1

    for fx {
        PrintColorText(color.FgHiYellow, "\nВаши посты:\n\n")
        PrintColorText(color.FgHiMagenta, strings.ToUpper(myPost[j].Title))
        PrintColorText(color.FgHiCyan, "\n" + myPost[j].DateTime.Format("02.01.2006 15:04"))
        fmt.Print("\n\n" + myPost[j].Text)
        fmt.Print("\n\n[0] - Вернуться в профиль\n[1] - Прошлый пост\n[2] - Следующий пост")
        var input string
        fmt.Scan(&input)
        if input == "1" {
            if j <= 0 {
                j = len(myPost) - 1
            }else {
                j--
            }

        }else if (input == "2")  {
            if j >= len(myPost) - 1 {
                j = 0
            }else {
                j++
            }
        }else {
            fx = false
            break
        }
        
    }
}

