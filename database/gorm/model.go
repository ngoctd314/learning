package main

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AccountUsernameModel struct {
	AccountID   string `gorm:"column:account_id"`
	Username    string `gorm:"column:username;primaryKey"`
	AccountType string `gorm:"column:account_type;primaryKey"`
}

func (AccountUsernameModel) TableName() string {
	return "account_usernames"
}

func (u AccountUsernameModel) BeforeCreate(tx *gorm.DB) error {
	if len(strings.TrimSpace(u.Username)) == 0 {
		return errors.New("invalid username")
	}
	if len(strings.TrimSpace(u.AccountType)) == 0 {
		return errors.New("invalid account_type")
	}

	return nil
}

type IAMClientModel struct {
	ClientID        string `gorm:"column:client_id"`
	ClientSecret    string `gorm:"column:client_secret"`
	ClientType      string `gorm:"column:client_type"`
	Method          string `gorm:"column:method"`
	Name            string `gorm:"column:name"`
	Description     string `gorm:"column:description"`
	PersonInCharge  string
	RedirectURIs    string
	AutoConsent     int
	Status          int
	Visible         int
	PasswordlessSms int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       string
}

func (IAMClientModel) TableName() string {
	return "iam_clients"
}
