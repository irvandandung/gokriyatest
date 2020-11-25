package repositorys

import(
	"github.com/go-pg/pg"
	"github.com/irvandandung/gokriyatest/models"
	"github.com/go-pg/pg/orm"
)

func GetAllDataUser(db *pg.DB, users []models.Users, limit int) ([]models.Users, error){
	err := db.Model(&users).ColumnExpr("users.id").ColumnExpr("users.data").ColumnExpr("roles.id AS roles__id").ColumnExpr("roles.data AS roles__data").
	Join("LEFT JOIN roles AS roles ON users.role_id = roles.id").Limit(limit).Select()
    return users, err
}


type Users struct{
	Id string  `json:"id" pg:"id,pk"`
	Data models.User `json:"data" pg:"data"`
	Role_id string `json:"role_id" pg:"role_id, fk"`
}

func InsertDataUser(db *pg.DB, user Users) (orm.Result, error){
	res, err := db.Model(&user).Insert()
    return res, err
}

func UpdateDataUser(db *pg.DB, user Users) (orm.Result, error){
	res, err := db.Model(&user).WherePK().Update()
    return res, err
}

func DeleteDataUser(db *pg.DB, id string) (orm.Result, error){
	var user Users
	res, err := db.Model(&user).Where("id = ?", id).Delete()
	return res, err
}