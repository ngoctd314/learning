package main

import "github.com/dchest/captcha"

// Captcha interface
type Captcha interface {
	Generate() ([]byte, error)
	Verify(id string, solution string) bool
}

type numcaptcha struct {
}

func newNumCaptcha(store captcha.Store) numcaptcha {
	captcha.SetCustomStore(store)

	return numcaptcha{}
}

func (n numcaptcha) Generate() ([]byte, error) {
	id := captcha.New()
	return []byte(id), nil
}

func (n numcaptcha) Verify(id string, solution string) bool {
	return captcha.VerifyString(id, solution)
}
