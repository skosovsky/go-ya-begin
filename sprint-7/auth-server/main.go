package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
)

const (
	// Объявляем привилегии нашей системы
	ReadPermission  = "read"
	WritePermission = "write"

	// Объявляем роли нашей системы
	AdminRole = "admin"
	UserRole  = "user"
)

var (
	// Связка роль — привилегии
	rolePermissions = map[string][]string{
		AdminRole: {ReadPermission, WritePermission},
		UserRole:  {ReadPermission},
	}
)

var (
	// Связка пользователь — роль
	userRoles = map[string][]string{
		"Alice2000": {UserRole},
		"Mike123":   {AdminRole},
	}
)

var (
	// Связка пользователь - пароль
	userPassword = map[string]string{
		"Alice2000": "17f80754644d33ac685b0842a402229adbb43fc9312f7bdf36ba24237a1f1ffb",
		"Mike123":   "d9803029af942f5fd8441a1e3649a29d1fc3251f011231a32a979c972f7d3178",
	}
)

// verifyUser — функция, которая выполняет аутентификацию и авторизацию пользователя.
// user — логин пользователя, pass — пароль, permission — необходимая привилегия.
// если пользователь ввел правильные данные, и у него есть необходимая привилегия — возвращаем true, иначе — false
func verifyUser(user, pass string, permission string) bool {
	// Процесс аутентификации начался.
	// получаем хеш пароля
	hashedPassword := sha256.Sum256([]byte(pass))
	hashStringPassword := hex.EncodeToString(hashedPassword[:])
	fmt.Println(hashStringPassword)
	// проверяем введенные данные
	storedPassword, ok := userPassword[user]
	// если пользователя нет в userPassword, то пользователя с таким логином не существует
	if !ok {
		return false
	}
	// если хеши паролей не совпадают — значит пароли не совпадают
	if hashStringPassword != storedPassword {
		return false
	}
	// Процесс аутентификации закончился.

	// Процесс авторизации начался.
	// итерируемся по всем привилегиям пользователя и ищем нужную нам
	for _, roles := range userRoles[user] {
		for _, storedPermission := range rolePermissions[roles] {
			if permission == storedPermission {
				return true
			}
		}
	}
	// если дошли до сюда, значит у пользователя нет нужной привилегии
	return false
	// Процесс авторизации закончился.
}

func main() {
	address := ":8080"

	mux := http.NewServeMux()
	// регистрируем endpoint /secret/
	mux.HandleFunc("/secret/", func(w http.ResponseWriter, req *http.Request) {
		// получаем данные пользователя
		user, pass, ok := req.BasicAuth()
		// проверяем доступы
		if ok && verifyUser(user, pass, ReadPermission) {
			fmt.Fprintf(w, "This is your secret: Hello world\n")
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})

	// Server — это структура для создания http-сервера, в которой можно более тонко его настроить.
	// В данном случае настроены два поля:
	// Addr — адрес, который будет прослушиваться, здесь ":8080" порт
	// Handler — маршрутизатор с зарегистрированными хендлерами, выше определён как mux.
	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s\n", address)
	// Для созданного и настроенного сервера вызываем ListenAndServe().
	err := srv.ListenAndServe()
	if err != nil {
		err = fmt.Errorf("failed to listen and serve: %w", err)
		fmt.Println(err)
	}
}
