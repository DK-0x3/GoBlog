package handler

import (
	"GoBlog/internal/database"
	"GoBlog/internal/database/models"
	middlewares "GoBlog/internal/middleWares"
	"errors"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var iUser database.LocalUser

func Entrance(email string, password string) (models.User, error) {
	users := iUser.GetUsers()

	for _, user := range *users {
		if user.Email == email {
			pas := string(middlewares.PasswordHash(password))
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pas)); err != nil {
				return user, nil
			}else {
				return models.User{}, errors.New("Неверный пароль")
			}
		}
	}
	return models.User{}, errors.New("Пользователь не найден")
}

func Registration(email string, name string) (string, bool) {
	users := iUser.GetUsers()
	nameValid := CheckValidName(name, users)
	emailValid := CheckValidEmail(email)

	
	if nameValid == "" {
		if emailValid == "" {

			passwd := generatePassword(10)
			Hashpasswd := string(middlewares.PasswordHash(passwd))

			newUser := models.User {
				Name: strings.ToLower(name),
				Email: email,
				Password: Hashpasswd,
			}
			*users = append(*users, newUser)
			return passwd, true

		}else {
			return emailValid, false
		}
		
	}else {
		return nameValid, false
	}
}

func generatePassword(length int) string {
	// символы, которые могут быть использованы в пароле
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*()_+-=:?/"
	password := make([]byte, length)
	rand.Seed(time.Now().UnixNano())

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func CheckValidName(name string, users *[]models.User) (string) {

	if len(name) >= 4 {
		for _, char := range name {
			re := regexp.MustCompile(`[[:punct:]]`)
 			if re.MatchString(string(char)) {
 				return "Никнейм не должен содержать спец символы"
 			}
		}
		for _, user := range *users {
			if user.Name == strings.ToLower(name) {
				return "Такой никнейм уже занят, попробуйте другой"
			}
		}
		return ""
	}else {
		return "Длинна никнейма минимум 4 символа"
	}
}

func CheckValidEmail(email string) (string) {
	validS := false
	validT := false
	users := iUser.GetUsers()

	for _, user := range *users {
		if user.Email == strings.ToLower(email) {
			return "Такая почта уже зарегистрирована!"
		}
	}

	for _, char := range email {
		if char == '@' {
			validS = true
		}
		if char == '.' {
			validT = true
		}
	}
	if validS && validT {
		return ""
	}else {
		return "Почта имеет неверный формат!"
	}
}