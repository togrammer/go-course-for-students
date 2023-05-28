package userrepo

import (
	"context"
	"testing"
)

func TestAddUser(t *testing.T) {
	repo := New()
	ctx := context.Background()

	user := repo.AddUser(ctx, "test", "test@example.com")

	if user.ID != 0 {
		t.Errorf("Expected ID to be 0, but got %v", user.ID)
	}

	if user.Nickname != "test" {
		t.Errorf("Expected Nickname to be 'test', but got %v", user.Nickname)
	}

	if user.Email != "test@example.com" {
		t.Errorf("Expected Email to be 'test@example.com', but got %v", user.Email)
	}
}

func TestFindUser(t *testing.T) {
	repo := New()
	ctx := context.Background()

	user := repo.AddUser(ctx, "test", "test@example.com")

	foundUser, err := repo.FindUser(ctx, user.ID)

	if err != nil {
		t.Errorf("Expected err to be nil, but got %v", err)
	}

	if foundUser.ID != user.ID {
		t.Errorf("Expected ID to be %v, but got %v", user.ID, foundUser.ID)
	}

	if foundUser.Nickname != user.Nickname {
		t.Errorf("Expected Nickname to be %v, but got %v", user.Nickname, foundUser.Nickname)
	}

	if foundUser.Email != user.Email {
		t.Errorf("Expected Email to be %v, but got %v", user.Email, foundUser.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	repo := New()
	ctx := context.Background()

	user := repo.AddUser(ctx, "test", "test@example.com")

	repo.UpdateUser(ctx, user.ID, "updated", "updated@example.com")

	foundUser, err := repo.FindUser(ctx, user.ID)

	if err != nil {
		t.Errorf("Expected err to be nil, but got %v", err)
	}

	if foundUser.Nickname != "updated" {
		t.Errorf("Expected Nickname to be 'updated', but got %v", foundUser.Nickname)
	}

	if foundUser.Email != "updated@example.com" {
		t.Errorf("Expected Email to be 'updated@example.com', but got %v", foundUser.Email)
	}
}
