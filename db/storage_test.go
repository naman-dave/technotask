package db

import "testing"

func TestAddUser(t *testing.T) {

	user := User{
		ID:        1,
		Username:  "username",
		Password:  "password",
		Firstname: "firstname",
		Lastname:  "lastname",
	}

	err := AddUser(1, user)
	if err != nil {
		t.Error("user should be added.", err)
	}

	user2 := User{
		ID:        1,
		Username:  "username",
		Password:  "password",
		Firstname: "firstname",
		Lastname:  "lastname",
	}

	err = AddUser(1, user2)
	if err == nil {
		t.Error("user should not be added because it already exists.")
	}

	user3 := User{
		ID:        2,
		Username:  "username",
		Password:  "password",
		Firstname: "firstname",
		Lastname:  "lastname",
	}

	err = AddUser(1, user3)
	if err == nil {
		t.Error("user should not be added because data mismatch")
	}

}

func TestDeleteUser(t *testing.T) {

	user := User{
		ID:        1,
		Username:  "username",
		Password:  "password",
		Firstname: "firstname",
		Lastname:  "lastname",
	}

	_ = AddUser(1, user)

	err := DeleteUser(1)
	if err != nil {
		t.Error("user should be deleted.", err)
	}

	err = DeleteUser(1)
	if err == nil {
		t.Error("user should not be deleted because it does not exist.")
	}
}

func TestEditUser(t *testing.T) {

	user := User{
		ID:        1,
		Username:  "username",
		Password:  "password",
		Firstname: "firstname",
		Lastname:  "lastname",
	}
	_ = AddUser(1, user)

	user.Firstname = "updated firstname"

	err := EditUser(1, user)
	if err != nil {
		t.Error("user should edited.", err)
	}

	output, err := GetUser(1)
	if err != nil {
		t.Error("user should be fetched.", err)
	}

	if output.Firstname != "updated firstname" {
		t.Errorf("user should be updated. expected %s got %s", "updated firstname", output.Firstname)
	}

	_ = DeleteUser(1)
	err = EditUser(1, user)
	if err == nil {
		t.Error("user should not be edited because it does not exist.")
	}
}
