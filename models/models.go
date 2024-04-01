package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/volatiletech/null"
)

type Organization struct {
	ID         int64     `json:"id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Website    string    `json:"website"`
	Sector     string    `json:"sector"`
	Status     string    `json:"status"`
	IsFinal    bool      `json:"isfinal"`
	IsArchived bool      `json:"isArchived"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Department struct {
	ID         int64     `json:"id"`
	Code       string    `json:"code"`
	OrgID      int64     `json:"orgID"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	IsFinal    bool      `json:"isfinal"`
	IsArchived bool      `json:"isArchived"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Role struct {
	ID           int64     `json:"id"`
	Code         string    `json:"code"`
	OrgID        int64     `json:"orgID"`
	DepartID     int64     `json:"departID"`
	Name         string    `json:"name"`
	Permissions  []string  `json:"permissions"`
	IsManagement bool      `json:"ismanagement"`
	Status       string    `json:"status"`
	IsFinal      bool      `json:"isfinal"`
	IsArchived   bool      `json:"isArchived"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type User struct {
	ID           int64     `json:"id"`
	Code         string    `json:"code"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"passwordHash"`
	IsAdmin      bool      `json:"isAdmin"`
	OrgID        int64     `json:"orgID"`
	RoleID       int64     `json:"roleID"`
	Status       string    `json:"status"`
	IsFinal      bool      `json:"isfinal"`
	IsArchived   bool      `json:"isArchived"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Login struct {
	Email        null.String `json:"email"`
	Phone        null.String `json:"phone"`
	OTP          string      `json:"otp"`
	PasswordHash string      `json:"passwordHash"`
}

type Otp struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userID"`
	Token     string    `json:"token"`
	IsValid   bool      `json:"isValid"`
	ExpiresAt time.Time `json:"expireAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Authtoken struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userID"`
	Token     uuid.UUID `json:"token"`
	IsValid   bool      `json:"isValid"`
	ExpireAt  time.Time `json:"expireAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Author struct {
	UserID int64     `json:"userID"`
	Name   string    `json:"name"`
	Token  uuid.UUID `json:"token"`
}

type Auther struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	IsAdmin      bool      `json:"IsAdmin"`
	OrgID        int64     `json:"orgID"`
	RoleID       int64     `json:"roleID"`
	SessionToken uuid.UUID `json:"sessionToken"`
}

type OTPRequest struct {
	Email null.String `json:"email"`
	Phone null.String `json:"phone"`
}

type Lead struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"userID"`
	FirstName  string    `json:"firstname"`
	LastName   string    `json:"lastname"`
	Address    string    `json:"address"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Occupation string    `json:"occupation"`
	Company    string    `json:"company"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
