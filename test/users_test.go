package test

import (
	"fmt"
	"testing"

	"github.com/yaska1706/AfyaPevu/internal/pkg/config"
	"github.com/yaska1706/AfyaPevu/internal/pkg/db"
	models "github.com/yaska1706/AfyaPevu/internal/pkg/models/users"
	"github.com/yaska1706/AfyaPevu/internal/pkg/persistence"
)

var userTest models.User

func Setup() {
	config.Setup("./config.yml")
	db.SetupDB()
	db.GetDB().Exec("DELETE FROM users")
}

func TestAddUser(t *testing.T) {
	Setup()
	user := models.User{
		Firstname: "Antonio",
		Lastname:  "Paya",
		Username:  "antonio",
		Hash:      "hash",
		Role:      models.UserRole{RoleName: "user"},
	}
	s := persistence.GetUserRepository()
	if err := s.Add(&user); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	userTest = user
}

func TestGetAllUsers(t *testing.T) {
	s := persistence.GetUserRepository()
	if _, err := s.All(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	db.SetupDB()
	db.SetupDB()
	s := persistence.GetUserRepository()
	if _, err := s.Get(fmt.Sprint(userTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
