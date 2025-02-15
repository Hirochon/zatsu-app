package domain

import (
	"time"

	"github.com/google/uuid"
)

// User はアプリケーションのユーザーを表す構造体です
type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser は新しいUserインスタンスを作成します
func NewUser(email, name string) *User {
	now := time.Now()
	return &User{
		ID:        generateID(), // 実際の実装では適切なID生成ロジックが必要です
		Email:     email,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateName はユーザーの名前を更新します
func (u *User) UpdateName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

// UpdateEmail はユーザーのメールアドレスを更新します
func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

// generateID はユーザーIDを生成します
func generateID() string {
	return uuid.New().String()
}
