package domain

import (
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	email := "test@example.com"
	name := "テストユーザー"

	user := NewUser(email, name)

	if user.Email != email {
		t.Errorf("期待したメールアドレス %v, 実際の値: %v", email, user.Email)
	}

	if user.Name != name {
		t.Errorf("期待した名前 %v, 実際の値: %v", name, user.Name)
	}

	if user.ID == "" {
		t.Error("IDが生成されていません")
	}

	if user.CreatedAt.IsZero() {
		t.Error("CreatedAtが設定されていません")
	}
}

func TestUpdateName(t *testing.T) {
	user := NewUser("test@example.com", "古い名前")
	oldUpdatedAt := user.UpdatedAt
	time.Sleep(time.Millisecond) // UpdatedAtの変更を確認するため

	newName := "新しい名前"
	user.UpdateName(newName)

	if user.Name != newName {
		t.Errorf("期待した名前 %v, 実際の値: %v", newName, user.Name)
	}

	if !user.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAtが更新されていません")
	}
}

func TestUpdateEmail(t *testing.T) {
	user := NewUser("old@example.com", "テストユーザー")
	oldUpdatedAt := user.UpdatedAt
	time.Sleep(time.Millisecond)

	newEmail := "new@example.com"
	user.UpdateEmail(newEmail)

	if user.Email != newEmail {
		t.Errorf("期待したメールアドレス %v, 実際の値: %v", newEmail, user.Email)
	}

	if !user.UpdatedAt.After(oldUpdatedAt) {
		t.Error("UpdatedAtが更新されていません")
	}
} 