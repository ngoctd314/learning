package main

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) CreateAccountUsername(u AccountUsernameModel) error {
	if err := r.db.Create(u).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) CreateAccountUsernameBatch(listUserName []AccountUsernameModel) error {
	if err := r.db.Create(listUserName).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) CreateAccountUsernameWithSelect(listUsername []AccountUsernameModel) error {
	if err := r.db.Select("username", "account_type", "account_id").Create(listUsername).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) BatchedInsertAccountUsername(listUserName []AccountUsernameModel, batchedSize int) error {
	if err := r.db.CreateInBatches(listUserName, batchedSize).Error; err != nil {
		return err
	}

	return nil
}

func (r repository) FirstUsername(rs *AccountUsernameModel) (*AccountUsernameModel, error) {
	if err := r.db.Model(rs).First(&rs).Error; err != nil {
		return nil, err
	}

	return rs, nil
}

func (r repository) TakeUsername(rs *AccountUsernameModel) (*AccountUsernameModel, error) {
	if err := r.db.Take(rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (r repository) LastUsername() (*AccountUsernameModel, error) {
	var rs AccountUsernameModel
	if err := r.db.Model(rs).Last(&rs).Error; err != nil {
		return nil, err
	}
	return &rs, nil
}

func (r repository) FindIAMClient() ([]IAMClientModel, error) {
	var rs []IAMClientModel

	if err := r.db.Limit(10).Find(&rs).Error; err != nil {
		return nil, err
	}

	return rs, nil
}
