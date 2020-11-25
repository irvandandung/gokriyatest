package main

import(
	"log"
	"github.com/irvandandung/gokriyatest/config"
	"github.com/gorilla/mux"
	"net/http"
)

func RunApp(){
	log.Println("init configuration...")
	pg := config.DbConn()
	r := mux.NewRouter()
	config.Routes(pg, r)
	port:= ":1234"
	log.Println("Starting server at "+port)
	log.Fatal(http.ListenAndServe(port, r))

	// status.Is_active = true
	// user.Email = "irvandandung1@gmail.com"
	// user.Password = "$2a$14$.EFXyVlcYNF0jDAf3sApu.lT.MrJB7a176P.Lj9W1snB2NBpB28Z6"
	// user.Username = "irvan"
	// user.Status = status

	// userInsert := repositorys.Users{
	// 	Id : "36cf2aec-2ea4-11eb-adc1-0242ac120002",
	// 	Data : user,
	// 	Role_id : "3094d252-2e4b-11eb-adc1-0242ac120002",
	// }
	// var users []models.Users
	// users = repositorys.GetAllDataUser(db, users, 5)
	// log.Println(userInsert)
	// res, err := repositorys.UpdateDataUser(db, userInsert)
	// res, err := repositorys.DeleteDataUser(db, "36cf2aec-2ea4-11eb-adc1-0242ac120002")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(res.RowsAffected())
	// log.Println(users)
	// log.Println(res.RowsAffected())
}