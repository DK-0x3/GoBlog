package services

import (
	"GoBlog/internal/database"
	"GoBlog/internal/handler"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// type BeginApp interface {
// 	StartProgram(StartTitle string)
// 	AuthorizationProfile()
// 	UserProfile()
// 	RegistrationProfile()
// }

func StartProgram(StartTitle string) {
	database.InitDB()
	for {
		ClearScreen()
		iActiveUser.DelActiveUser()
		Title := StartTitle + "\n\n[1] - Войти\n[2] - Зарегистрироваться\n[3] - Просмотр недавних постов пользователей\n[0] - завершить сеанс\nВведи команду: "
		PrintColorText(color.FgBlue, Title)

		var action string
		fmt.Scan(&action)

		switch action {
		case "0":
			ClearScreen()
			fmt.Println("Завершение программы...")
			os.Exit(0)
		case "1":
			AuthorizationProfile()
		case "2":
			RegistrationProfile()
		case "3":
			ClearScreen()
			ViewAllPosts()
		default:
			fmt.Print("Неверный ввод, такой команды не найдено")
		}
	}
}

func AuthorizationProfile() {
	for {
		ClearScreen()
		var email string
		var password string
		PrintColorText(color.FgGreen, "Вход\n\n*Вернуться назад - [0]\n\n")
		fmt.Print("Введи Email: ")
		fmt.Scan(&email)
		if email == "0" {
			break
		}
		fmt.Print("Введи Пароль: ")
		fmt.Scan(&password)
		if password == "0" {
			break
		}
		activUser, err := handler.Entrance(email, password)

		if err != nil {
			var input string
			PrintColorText(color.FgRed, err.Error())
			fmt.Print("\n[1] - Попробовать снова\n[0] - Вернуться на главную\nВаш выбор: ")
			fmt.Scan(&input)
			if input == "0" {
				break
			} else if input == "1" {
				continue
			} else {
				break
			}
		}
		
		iActiveUser.PutActiveUser(activUser)

		UserProfile()
	}
}

func UserProfile() {
	for {
		ClearScreen()
		PrintColorText(color.FgGreen, "Добро пожаловать " + iActiveUser.GetActiveUser().Name)
		fmt.Print("\n[0] - Выход из профиля\n[1] - Просмотр и редактирование постов\n[2] - Создать новый пост\n[3] - Просмотр ленты постов\n")
		
		var inputProfile string
		fmt.Scan(&inputProfile)
		if inputProfile == "0" {
			break
		} else if inputProfile == "1" {
			ViewMyPosts()
		} else if inputProfile == "2" {
			ClearScreen()
			AddPosts()
		}else if inputProfile == "3" {
			ClearScreen()
			ViewAllPosts()
		}
	}
}

func RegistrationProfile() {
	for {
		ClearScreen()
		var email string
		var name string
		PrintColorText(color.FgGreen, "Регистрация\n\n*Вернуться назад - [0]\n\n")
		fmt.Print("Введи Email: ")
		fmt.Scan(&email)
		if email == "0" {
			break
		}

		fmt.Print("Введи Name: ")
		fmt.Scan(&name)
		if name == "0" {
			break
		}

		result, registrValid := handler.Registration(email, name)
		if registrValid {
			PrintColorText(color.FgGreen, "Успешная регистрация!\nОсталось войти в профиль\n\n")
			fmt.Printf("Ваш пароль: %s\n\nВведи что угодно для перехода на главную...\n", result)
			fmt.Scan(&result)
			break
		} else {
			var input string
			PrintColorText(color.FgRed, "Ошибка регистрации: "+result)
			fmt.Print("\n[1] - Попробовать снова\n[0] - Вернуться на главную\nВаш выбор: ")
			fmt.Scan(&input)
			if input == "0" {
				break
			} else if input == "1" {
				continue
			} else {
				break
			}
		}

	}
}

