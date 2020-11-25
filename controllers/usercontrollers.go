package controllers

import(
	"github.com/go-pg/pg"
	"net/http"
	"log"
	"github.com/irvandandung/gokriyatest/models"
	"github.com/irvandandung/gokriyatest/repositorys"
	"github.com/irvandandung/gokriyatest/helpers"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
)

type Db struct{
	Pg *pg.DB
}

type MyClaims struct {
    jwt.StandardClaims
    User models.User
}

func (db *Db) UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pg := db.Pg
	if r.Method != "POST" {
        http.Error(w, "Unsupported http method", http.StatusBadRequest)
        return
    }

    username, password, ok := r.BasicAuth()
    if !ok {
        http.Error(w, "Invalid username or password", http.StatusBadRequest)
        return
    }

    var user models.Users

    res, err := repositorys.GetDataUserByUsername(pg, user, username)
    if err != nil{
    	http.Error(w, "User not found!", http.StatusUnauthorized)
        return
    }

    check := helpers.CheckPasswordHash(res.Data.Password, password)
    if check != true {
    	http.Error(w, "Wrong password!", http.StatusUnauthorized)
        return
    }

    claims := MyClaims{
        StandardClaims: jwt.StandardClaims{
            Issuer:    APPLICATION_NAME,
            ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
        },
        User : res.Data,
    }
    log.Println(claims)

    token := jwt.NewWithClaims(
        JWT_SIGNING_METHOD,
        claims,
    )

    signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"token": signedToken})
}

func (db *Db) GetAllDataUser(w http.ResponseWriter, r *http.Request) {
	
}

func (db *Db) GetDataUser(w http.ResponseWriter, r *http.Request) {
	
}

func (db *Db) InsertDataUser(w http.ResponseWriter, r *http.Request) {
	
}

func (db *Db) UpdateDataUser(w http.ResponseWriter, r *http.Request) {
	
}

func (db *Db) DeleteDataUser(w http.ResponseWriter, r *http.Request) {
	
}