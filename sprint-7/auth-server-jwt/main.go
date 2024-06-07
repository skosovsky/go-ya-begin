package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret = []byte("gBElG5NThZSye") //nolint:gochecknoglobals // it's learning code
)

const (
	ReadPermission = "read" // Объявляем привилегии нашей системы
	UserRole       = "user" // Объявляем роли нашей системы
)

var (
	// Связка роль — привилегии.
	rolePermissions = map[string][]string{ //nolint:gochecknoglobals // it's learning code
		UserRole: {ReadPermission},
	}
)

var (
	// Связка пользователь — роль.
	userRoles = map[string][]string{ //nolint:gochecknoglobals // it's learning code
		"Alice2000": {UserRole},
	}
)

// verifyUser — функция, которая выполняет аутентификацию и авторизацию пользователя.
// Token — JWT пользователя, если пользователь ввел правильные данные и у него есть необходимая привилегия - возвращаем true, иначе - false.
func verifyUser(token string, permission string) bool {
	jwtToken, err := jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		log.Printf("Failed to parse token: %s\n", err)
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
		log.Printf("failed to sign jwt: %s\n", err)
	}
	log.Println("Result token: " + signedToken)
}

func main() {
	getToken()

	address := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/secret/", func(w http.ResponseWriter, req *http.Request) {
		// получаем http header вида 'Bearer {jwt}'
		authHeaderValue := req.Header.Get("Authorization")
		// проверяем доступы
		if authHeaderValue != "" { //nolint:nestif // it's learning code
			bearerToken := strings.Split(authHeaderValue, " ")
			if len(bearerToken) == 2 { //nolint:mnd // it's learning code
				if verifyUser(bearerToken[1], ReadPermission) {
					partKey := strings.Split(bearerToken[1], ".")[0]
					_, err := fmt.Fprintf(w, "This is your secret: %s\n", partKey)
					if err != nil {
						return
					}
					return
				}
			}
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})

	srv := &http.Server{ //nolint:gosec // it's learning code
		Addr:    address,
		Handler: mux,
	}

	log.Printf("Starting server on %s\n", address)
	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("failed to listen and serve: %s\n", err)
	}
}
