package data

import (
	"errors"

	"github.com/muling3/go-fiber-jwt/models"
)

var usersList = []models.User{
	{Id: 1, Username: "George Boole", Email: "george.boole@gmail.com", Password: "12345"},
	{Id: 2, Username: "Roy Fielding", Email: "roy.fielding@yahoo.com", Password: "1f345"},
	{Id: 3, Username: "James Gosling", Email: "jamesG@gmail.com", Password: "1235"},
	{Id: 4, Username: "Linus Torvalds", Email: "linus.torvalds@yahoo.com", Password: "1345"},
	{Id: 5, Username: "Mulinge Muli", Email: "muli.mulinge@gmail.com", Password: "2345"},
}

func AddUser(user models.User) models.User {
	user.Id = len(usersList) + 1
	usersList = append(usersList, user)
	return user
}

func GetUsers() []models.User {
	return usersList
}

func GetUser(id int) models.User {
	for _, u := range usersList {
		if u.Id == id {
			return u
		}
	}

	return models.User{}
}

func RemoveUser(id int) error {
	for i, u := range usersList {
		if u.Id == id {
			usersList = append(usersList[:i], usersList[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found")
}

func Login(username, password string) (models.User, error) {
	for _, u := range usersList {
		if u.Username == username && u.Password == password {
			return u, nil
		}
	}

	return models.User{}, errors.New("invalid login credentials")
}
