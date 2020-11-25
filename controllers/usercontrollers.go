package controllers

import(
	"github.com/go-pg/pg"
	"net/http"
	"log"
	"strconv"
	"github.com/irvandandung/gokriyatest/models"
	"github.com/irvandandung/gokriyatest/repositorys"
	"github.com/irvandandung/gokriyatest/helpers"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"github.com/go-pg/pg/orm"
)

type Db struct{
	Pg *pg.DB
}

type MyClaims struct {
    jwt.StandardClaims
    User models.Users
}

func (db *Db) UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pg := db.Pg
	if r.Method != "POST" {
        http.Error(w, "Unsupported http method!", http.StatusBadRequest)
        return
    }

    username, password, ok := r.BasicAuth()
    if !ok {
        http.Error(w, "Invalid username or password!", http.StatusBadRequest)
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
        User : res,
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
	w.Header().Set("Content-Type", "application/json")

	pg := db.Pg
	var user []models.Users
	var err error
	if r.Method != "GET" {
        http.Error(w, "Unsupported http method!", http.StatusBadRequest)
        return
    }

    keys, ok := r.URL.Query()["page"]
    if !ok {
    	http.Error(w, "Please add parameter page!", http.StatusBadRequest)
    	return 
    }

    page, _ := strconv.Atoi(keys[0])

    limit := 5
    offset := (page - 1) * 5

    user, err = repositorys.GetAllDataUser(pg, user, limit, offset)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

    var data []interface{}
    for _, val := range user {
    	userdata := map[string]interface{}{
	    	"Username" : val.Data.Username,
	    	"Email" : val.Data.Email,
	    	"Status" : val.Data.Status,
    	}
    	data = append(data, userdata)
    }
    log.Println(data)

    json.NewEncoder(w).Encode(data)
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userInfo := userInfo(r.Context().Value("userInfo").(jwt.MapClaims))
	json.NewEncoder(w).Encode(userInfo)
}


func (db *Db) InsertDataUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := userInfo(r.Context().Value("userInfo").(jwt.MapClaims))
    pg := db.Pg
	var user models.User
	var users models.Users
	var err error
	var uuid string
	var res orm.Result
	var response map[string]interface{}

	if r.Method != "POST" {
        http.Error(w, "Unsupported http method!", http.StatusBadRequest)
        return
    }

    role_name := userInfo["Role Name"].(string)
    if role_name != "Admin"{
    	http.Error(w, "Sorry, You can't access this request because you're not admin!", http.StatusUnauthorized)
        return
    }

    err = json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _ , err = repositorys.GetDataUserByUsername(pg, users, user.Username)
    if err == nil{
    	http.Error(w, "Sorry, the username has been taken!", http.StatusUnauthorized)
        return
    }

    uuid,err = helpers.UuidGenerate()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    user.Password, err = helpers.HashPassword(user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    userInsert := repositorys.Users{
		Id : uuid,
		Data : user,
		Role_id : "d57bfbfe-4979-4809-a151-f6cd30de657b",
	}
	res, err = repositorys.InsertDataUser(pg, userInsert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	if res.RowsAffected() == 1 {
		response = map[string]interface{}{
			"status" : res.RowsAffected(),
			"message" : "Insert successful!",
		}
	}else{
		response = map[string]interface{}{
			"status" : res.RowsAffected(),
			"message" : "Insert not was successful!",
		}
	}
	
	json.NewEncoder(w).Encode(response)
}

func (db *Db) UpdateDataUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pg := db.Pg
	userInfo := userInfo(r.Context().Value("userInfo").(jwt.MapClaims))
	var res orm.Result
	var user models.User
	var response map[string]interface{}
	var err error

	if r.Method != "POST" {
        http.Error(w, "Unsupported http method!", http.StatusBadRequest)
        return
    }

	role_name := userInfo["Role Name"].(string)
    if role_name != "Admin"{
    	http.Error(w, "Sorry, You can't access this request because you're not admin!", http.StatusUnauthorized)
        return
    }

	err = json.NewDecoder(r.Body).Decode(&user)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user.Password, err = helpers.HashPassword(user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    keys, ok := r.URL.Query()["id"]
    if !ok {
    	http.Error(w, "Please add parameter id!", http.StatusBadRequest)
    	return 
    }

    id := keys[0]

    userUpdate := repositorys.Users{
    	Id : id,
    	Data : user,
    	Role_id : "d57bfbfe-4979-4809-a151-f6cd30de657b",
    }

    res, err = repositorys.UpdateDataUser(pg, userUpdate)
    if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	if res.RowsAffected() == 1 {
		response = map[string]interface{}{
			"status" : res.RowsAffected(),
			"message" : "Update successful!",
		}
	}else{
		response = map[string]interface{}{
			"status" : res.RowsAffected(),
			"message" : "Data id not found!",
		}
	}
	
	json.NewEncoder(w).Encode(response)
}

func (db *Db) DeleteDataUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pg := db.Pg
	userInfo := userInfo(r.Context().Value("userInfo").(jwt.MapClaims))
	var response map[string]interface{}

	role_name := userInfo["Role Name"].(string)
    if role_name != "Admin"{
    	http.Error(w, "Sorry, You can't access this request because you're not admin!", http.StatusUnauthorized)
        return
    }

	if r.Method != "DELETE" {
        http.Error(w, "Unsupported http method!", http.StatusBadRequest)
        return
    }    

    keys, ok := r.URL.Query()["id"]
    if !ok {
    	http.Error(w, "Please add parameter id!", http.StatusBadRequest)
    	return 
    }

    id := keys[0]
    log.Println(id)

    res, err := repositorys.DeleteDataUser(pg, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}

	if res.RowsAffected() == 1 {
		response = map[string]interface{}{
			"status" : res.RowsAffected(),
			"message" : "Delete successful!",
		}
	}else{
		response = map[string]interface{}{
			"status" : res.RowsAffected(),
			"message" : "Data id not found!",
		}
	}
	
	json.NewEncoder(w).Encode(response)
}


func userInfo(val jwt.MapClaims) map[string]interface{} {
	user := val["User"].(map[string]interface {})
	datauser := user["data"].(map[string]interface {})
	role := user["Roles"].(map[string]interface {})
	datarole := role["data"].(map[string]interface {})

	return map[string]interface{}{
		"User Id": user["id"],
		"Username": datauser["username"],
		"Email": datauser["email"],
		"Role Name": datarole["role_name"],
	}
}