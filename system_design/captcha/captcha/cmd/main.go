package main

import (
	"fmt"
	"time"

	"github.com/dchest/captcha"
)

func main() {
	store := captcha.NewMemoryStore(100, time.Minute)
	captcha.SetCustomStore(store)

	id := captcha.New()
	// fmt.Println(string(store.Get(id, true)))
	// fmt.Println(id, store, )
	fmt.Println(captcha.Verify(id, store.Get(id, false)))
}
