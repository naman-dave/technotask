package db

import "fmt"

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

var userStorage = map[int]User{}

func GetUsers() []User {
	users := []User{}

	for _, user := range userStorage {
		users = append(users, user)
	}

	return users
}

func GetUser(id int) (User, error) {

	if _, ok := userStorage[id]; !ok {
		return User{}, fmt.Errorf("User does not exists")
	}

	return userStorage[id], nil
}

func AddUser(id int, user User) error {

	if _, ok := userStorage[id]; ok {
		return fmt.Errorf("User already exists")
	}

	if id != user.ID {
		return fmt.Errorf("Data mismatch between id and user_id")
	}

	userStorage[id] = user

	return nil
}

func EditUser(id int, user User) error {

	if _, ok := userStorage[id]; !ok {
		return fmt.Errorf("User does not exists")
	}

	if id != user.ID {
		return fmt.Errorf("Data mismatch between id and user_id")
	}

	userStorage[id] = user

	return nil
}

func DeleteUser(id int) error {
	if _, ok := userStorage[id]; !ok {
		return fmt.Errorf("User does not exists")
	}

	delete(userStorage, id)

	return nil
}
