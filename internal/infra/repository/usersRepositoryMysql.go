package repository

import (
	"GolangAuthetication/internal/entity"
	"database/sql"
)

type UserRepositoryMysql struct {
	Db *sql.DB
}

func NewUserRepositoryMysql(Db *sql.DB) *UserRepositoryMysql {
	return &UserRepositoryMysql{Db: Db}
}

func (r *UserRepositoryMysql) GetUser() ([]*entity.User, error) {

	rows, err := r.Db.Query("Select id ,name, email, password from Users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
func (r UserRepositoryMysql) CreateUser(User *entity.User) error {
	_, err := r.Db.Exec("Insert into Users (id,name,email,password,create_time) values (?,?,?,?,?)", User.Id, User.Name, User.Email, User.Password, User.Create_time)
	if err != nil {
		return err
	}
	return nil
}
func (r UserRepositoryMysql) GetPerEmail(User *entity.User) (entity.User, error) {
	var user entity.User
	row := r.Db.QueryRow("Select id ,name, email, password from Users Where email = ?", User.Email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}
	return user, nil
}
