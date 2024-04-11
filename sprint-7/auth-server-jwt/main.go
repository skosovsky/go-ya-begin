package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret = []byte("gBElG5NThZSye")
)

const (
	// Объявляем привилегии нашей системы.
	ReadPermission = "read"

	// Объявляем роли нашей системы.
	UserRole = "user"
)

var (
	// Связка роль — привилегии.
	rolePermissions = map[string][]string{
		UserRole: {ReadPermission},
	}
)

var (
	// Связка пользователь — роль.
	userRoles = map[string][]string{
		"Alice2000": {UserRole},
	}
)

// verifyUser — функция, которая выполняет аутентификацию и авторизацию пользователя.
// token — JWT пользователя.
// если у пользователь ввел правильные данные,и у него есть необходимая привилегия - возвращаем true, иначе - false.
func verifyUser(token string, permission string) bool {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		fmt.Printf("Failed to parse token: %s\n", err)
		return false
	}
	if !jwtToken.Valid {
		return false
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	loginRaw, ok := claims["login"]
	if !ok {
		return false
	}

	login, ok := loginRaw.(string)
	if !ok {
		return false
	}

	for _, roles := range userRoles[login] {
		for _, storedPermission := range rolePermissions[roles] {
			if permission == storedPermission {
				return true
			}
		}
	}
	return false
}

func getToken() {
	// создаём payload
	claims := jwt.MapClaims{
		"login": "Alice2000",
		"roles": []string{"reader"},
	}

	// создаём jwt и указываем payload
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// получаем подписанный токен
	signedToken, err := jwtToken.SignedString(secret)
	if err != nil {
		fmt.Printf("failed to sign jwt: %s\n", err)
	}
	fmt.Println("Result token: " + signedToken)
}

func main() {
	getToken()

	address := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/secret/", func(w http.ResponseWriter, req *http.Request) {
		// получаем http header вида 'Bearer {jwt}'
		authHeaderValue := req.Header.Get("Authorization")
		// проверяем доступы
		if authHeaderValue != "" {
			bearerToken := strings.Split(authHeaderValue, " ")
			if len(bearerToken) == 2 {
				if verifyUser(bearerToken[1], ReadPermission) {
					partKey := strings.Split(bearerToken[1], ".")[0]
					fmt.Fprintf(w, "This is your secret: %s\n", partKey)
					return
				}
			}
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s\n", address)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to listen and serve: %s\n", err)
	}
}
