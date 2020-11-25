package controllers

import(
	"github.com/go-pg/pg"
	"net/http"
)

type Db struct{
	Pg *pg.DB
}

func (db *Db) UserLogin(w http.ResponseWriter, r *http.Request) {
	
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