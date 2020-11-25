package controllers

import (
	"net/http"
	"fmt"
    "context"
    "strings"
    "time"
    jwt "github.com/dgrijalva/jwt-go"
)

var APPLICATION_NAME = "Simple Api with Go & Pg"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("signature of my API")

func JwtAuth(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/user/login" {
            next.ServeHTTP(w, r)
            return
        }

        authorizationHeader := r.Header.Get("Authorization")
        if !strings.Contains(authorizationHeader, "Bearer") {
            http.Error(w, "Invalid token", http.StatusBadRequest)
            return
        }

        tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Signing method invalid")
            } else if method != JWT_SIGNING_METHOD {
                return nil, fmt.Errorf("Signing method invalid")
            }

            return JWT_SIGNATURE_KEY, nil
        })
        
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        ctx := context.WithValue(context.Background(), "userInfo", claims)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}