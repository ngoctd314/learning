package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// Store ...
type Store struct {
	db *sqlx.DB
}

// NewStore create new store
func NewStore(db *sqlx.DB) Store {
	return Store{db}
}

// Set sets the digits for the captcha id.
func (s Store) Set(id string, digits []byte) {
	challenge := Challenge{
		ID:           id,
		ClientID:     "test-clientid",
		Question:     "Type the numbers you see in the picture below",
		QuestionType: QuestionTypeNumber,
		Answer:       ConvByteNumbersToStringNumbers(digits),
		CreatedAt:    time.Now(),
		ExpiredAt:    time.Now().Add(time.Minute),
	}
	_, err := s.db.NamedExec(`INSERT INTO challenges (id, client_id, question, question_type, data, answer,  created_at, expired_at) 
	VALUES (:id, :client_id, :question, :question_type, :data, :answer,  :created_at, :expired_at)`, challenge)
	if err != nil {
		log.Println(err)
	}
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (s Store) Get(id string, clear bool) (digits []byte) {
	var challenge Challenge

	err := s.db.Get(&challenge, "SELECT answer FROM challenges WHERE id = ?", id)
	if err != nil {
	}
	digits = ConvStringNumbersToByteNumbers(challenge.Answer)
	fmt.Println("digits", digits, challenge.Answer)

	return
}

// QuestionType ...
type QuestionType int8

// QuestionType values
const (
	QuestionTypeNumber QuestionType = 1
)

// Challenge ...
type Challenge struct {
	ID           string       `db:"id"`
	ClientID     string       `db:"client_id"`
	Question     string       `db:"question"`
	QuestionType QuestionType `db:"question_type"`
	Data         string       `db:"data"`
	Answer       string       `db:"answer"`
	IsUsed       bool         `db:"is_used"`
	CreatedAt    time.Time    `db:"created_at"`
	ExpiredAt    time.Time    `db:"expired_at"`
}
