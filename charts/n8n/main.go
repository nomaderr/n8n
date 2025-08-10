package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "newpassword123" // замените на нужный пароль

	// Генерируем bcrypt-хэш с cost = 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	fmt.Println("bcrypt hash:", string(hash))
}
