package db

import (
	"testing"

	"github.com/nikhosagala/soapi/models"
)

func TestAdmin(t *testing.T) {
	CreateConnection()
	defer CloseConnection()
	admin1 := models.Admin{Username: "nikhosagala", Password: ""}

	_, err := CreateAdmin(&admin1)

	// check if the password empty
	if err == nil {
		t.Error("Should be error, admin should not created.")
	}

	admin1.Password = "sagala"

	adminSave, err2 := CreateAdmin(&admin1)

	if err2 != nil {
		t.Errorf("Error is %s", err2.Error())
	}

	// check if password is encrypted
	if adminSave.Password == "sagala" {
		t.Errorf("Password should encrypted. Password is %s", adminSave.Password)
	}

	// check if admin is created when username already exist
	admin2 := models.Admin{Username: "nikhosagala", Password: "sagala"}

	username, err3 := CreateAdmin(&admin2)

	if err3 == nil {
		t.Errorf("Admin should not be created. The username is %s", username.Username)
	}

	size := len(ReadAllAdmins(0, 0))

	if size == 0 {
		t.Errorf("Table admins should not be empty. Size %d", size)
	}

	adminRead := ReadAdmin(adminSave.ID)

	if adminRead.ID == 0 {
		t.Errorf("Admin should be found. ID is %d", adminRead.ID)
	}

	adminRead.Username = "nikhosagala"

	adminUpdate, err4 := UpdateAdmin(&adminRead)

	// check if updated username already exist
	if err4 == nil {
		t.Errorf("Username already exist, should be error, username is %s", adminUpdate.Username)
	}

	// check if updated password is encrypted
	if adminUpdate.Password == "word" {
		t.Errorf("Password should be encrypted, password is %s", adminUpdate.Password)
	}

	// check if password not updated
	if adminRead.Password != adminUpdate.Password {
		t.Errorf("Password should be equal, password is %s", adminRead.Password)
	}

	adminDeleted := DeleteAdmin(adminUpdate)

	if &adminDeleted.DeletedAt == nil {
		t.Error("Admin should be deleted.")
	}
}

func TestConnection(t *testing.T) {

}
