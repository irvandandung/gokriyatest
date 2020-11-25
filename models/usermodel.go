package models


type Roles struct{
	Id string  `json:"id" pg:"id,pk"`
	Data Role `json:"data" pg:"data"`
}

type Role struct{
	Role_name string `json:"role_name"`
	Description string `json:"description"`
}

type Users struct{
	tableName struct{} `pg:"select:users"`
	Id string  `json:"id" pg:"id,pk"`
	Data User `json:"data" pg:"data"`
	Role_id string `json:"role_id" pg:"role_id, fk"`
 	Roles *Roles `pg:"-"`
}

type User struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Status Status
}

type Status struct{
	Is_active bool `json:"is_active"`
}