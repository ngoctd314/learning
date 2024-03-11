package numcaptcha

import (
	"errors"
	"time"
)

// implements generation and verification of image CAPTCHAs

// A captcha solution is the sequence of digits 0-9 with the defined length
// An image representation is a PNG-encoded image with the solution printed on
// it in such a way that makes it hard for computers to solve it using OCR.

// Doesn't require external files or libraries to generate captcha representations; it is self-contained

// To make captchas one-time, the package includes a memory storage that stores captcha ids
// their solutions, and expiration time. Used captchas are removed from the store immediately
// after calling verify or verifyString, while unused captchas (user loaded a page with captcha, but
// didn't submit the form) are collected automatically after the predefined expiration time.
// Developers can also provide custom store (for example, which saves captcha ids and solutions in database)
// by implementing Store interface and registering the object with SetCustomStore

// Captchas are created by calling New, which returns the captcha id. Their representations, though
// are created on-the-fly by calling WriteImage or WriteAudio functions. Created representations are not
// stored anywhere, but subsequent calls to these functions with th same id will write the same captcha solution.

// Server provides an http.Handler which can serve large and audio representations of captchas automatically
// from the URL. It can also be used to reload captchas. Refer to Server function documentation for details
// or take a look at the example in "capexample" subdirectory.

const (
	// DefaultLen number of digits in captcha solution
	DefaultLen = 6
	// CollectNum the number of captchas created that triggers garbage collection used by default store
	CollectNum = 100
	// Expiration time of captchas used by default store
	Expiration = 10 * time.Minute
)

// all errors here
var (
	ErrNotFound       = errors.New("captcha: id not found")
	globalStore Store = NewMemoryStore(CollectNum, Expiration)
)

// SetCustomStore sets custom storage for captchas, replacing the default memory store.
// This function must be called before generating any captchas.
func SetCustomStore(s Store) {
	globalStore = s
}

// New creates a new captcha with the standard length, saves it in the internal
// storage and returns its id
func New() string {
	// XWo24cwFdBO2ovZZjK2S
	return NewLen(DefaultLen)
}

// NewLen ...
func NewLen(length int) (id string) {
	id = randomID()
	globalStore.Set(id, RandomDigits(length))
	return
}
