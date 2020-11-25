package config

import(
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"github.com/irvandandung/gokriyatest/controllers"
)

func Routes(pg *pg.DB, r *mux.Router){
	db := controllers.Db{Pg:pg}
	r.HandleFunc("/user/login", db.UserLogin)
	r.HandleFunc("/user/get-list-user", db.GetAllDataUser)
	r.HandleFunc("/user/get-user-data", controllers.GetUserData)
	r.HandleFunc("/user/insert-user", db.InsertDataUser)
	r.HandleFunc("/user/update-user", db.UpdateDataUser)
	r.HandleFunc("/user/delete-user", db.DeleteDataUser)

	//add middleware with jwt security
	r.Use(controllers.JwtAuth)
}