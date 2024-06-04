package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	googleSignin = "https://accounts.google.com"
)

func newChromedp() (context.Context, context.CancelFunc) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("remote-debugging-port", "9222"),
		chromedp.Flag("mute-audio", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	email := "//*[@id='user[email]']"
	passwd := "//*[@id='user[password]']"
	task := chromedp.Tasks{
		chromedp.Navigate("https://www.rachelsenglishacademy.com/enrollments"),
		chromedp.SendKeys(email, ""),
		chromedp.SendKeys(passwd, ""),
		chromedp.Click("//*[@id='sign-in']"),
		chromedp.Sleep(time.Second * 5),
		// link
		chromedp.Navigate("link"),
		chromedp.Sleep(time.Second * 5),
		// start
		chromedp.Click("//*[@id='w-vulcan-v2-42']/div[4]/div/div[3]/div/div/div/button"),
		// full screen
		// set cc
		chromedp.Sleep(time.Minute * 9),
	}

	if err := chromedp.Run(ctx, task); err != nil {
		fmt.Println(err)
	}

	return ctx, cancel
}

func main() {
	_, cancel := newChromedp()
	defer cancel()
}
