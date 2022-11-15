package db

import (
	"database/sql"
	"errors"

	"github.com/muling3/go-fiber-jwt/models"
)

func FetchUsers() ([]models.User, error){
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usersList []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		usersList = append(usersList, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return usersList, nil
}

func GetUserById(id int) (models.User, error) {
	stmt, err := DB.Prepare("SELECT id, username, email FROM users WHERE id = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var user models.User
	if err := stmt.QueryRow(id).Scan(&user.Id, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, errors.New("no User by given id")
		} else {
			return models.User{}, err
		}
	}
	return user, nil
}

func RegisterUser(user models.User) (int64, error) {
	stmt, err := DB.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	insertedId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return insertedId, nil
}

func Login(user models.User) (models.User, error) {
	stmt, err := DB.Prepare("SELECT * FROM users WHERE username = ? AND password = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	var u models.User
	err = stmt.QueryRow(user.Username, user.Password).Scan(&u.Id,&u.Username, &u.Email, &u.Password)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func RemoveUser(id int) error {
	stmt, err := DB.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	stmt.Exec(id)
	return nil
}

func UpdateUser(id int, user models.User) (models.User, error){
	//get the user by id
	stmt, err := DB.Prepare("UPDATE users SET username=?, email=?, password=? WHERE id = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password, id)
	if err != nil {
		return models.User{}, err
	}

	u, err := GetUserById(id)
	if err != nil{
		return models.User{}, err
	}

	return u, nil
}
