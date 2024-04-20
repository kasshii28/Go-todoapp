package models

import (
	"log"
	"time"
)

type User struct {
	ID int
	UUID string
	Name string
	Email string
	PassWord string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error){
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(
			cmd,
			createUUID(),
			u.Name,
			u.Email,
			Encrypt(u.PassWord), 
			time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//ユーザーの取得
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created-at ftom usrs where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

//ユーザーの更新
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}