package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/jpillora/opts"
)

var version = "0.0.0-src"

var config = struct {
	Portal      string        `opts:"help=sungrow portal, env=PORTAL"`
	Email       string        `opts:"help=sungrow username, env=EMAIL"`
	Pass        string        `opts:"help=sungrow email, env=PASSWORD"`
	Timeout     time.Duration `help:"overall timeout"`
	NoHeadless  bool          `help:"disable headless mode"`
	ShowNetwork bool          `help:"show network requests"`
	Debug       bool          `help:"show chromedp actions"`
	Screenshot  string        `help:"take a screenshot, and write it to this path"`
}{
	Portal:  `https://www.isolarcloud.com`,
	Timeout: 30 * time.Second,
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	opts.New(&config).Name(`sungrow-chromedp`).Version(version).Parse()
	if config.Email == "" || config.Pass == "" {
		log.Fatal("email and password are required")
	}

	// create a timeout
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	ctx, cancel = chromedp.NewExecAllocator(
		ctx,
		chromedp.Flag("headless", !config.NoHeadless),
		chromedp.Flag("remote-debugging-port", "9222"),
	)
	defer cancel()

	copts := []chromedp.ContextOption{}
	copts = append(copts, chromedp.WithLogf(log.Printf))
	copts = append(copts, chromedp.WithErrorf(log.Printf))
	// copts = append(copts, chromedp.WithDebugf(log.Printf))
	// if config.Debug {
	// }
	ctx, cancel = chromedp.NewContext(ctx, copts...)
	defer cancel()

	if config.ShowNetwork {
		chromedp.ListenTarget(ctx, func(event interface{}) {
			switch e := event.(type) {
			case *network.EventLoadingFinished:
				log.Printf("event: finished-loading: %s", e.RequestID)
			case *network.EventRequestWillBeSent:
				log.Printf("event: request: %s %s => #%s", e.Request.Method, e.Request.URL, e.RequestID)
			case *network.EventResponseReceived:
				log.Printf("event: response: #%s => %d", e.RequestID, e.Response.Status)
			}
		})
	}
	// result
	var kwh string
	var screenshot []byte
	// optionally save it to file
	if config.Screenshot != "" {
		defer func() {
			if len(screenshot) > 0 {
				os.WriteFile(config.Screenshot, screenshot, 0644)
				log.Printf("wrote screenshot to: %s", config.Screenshot)
			}
		}()
	}
	// navigate to a page, wait for an element, click
	err := chromedp.Run(ctx,
		chromedp.Navigate(config.Portal),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.WaitVisible(`#login-btn`),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.SendKeys(`#userAcct`, config.Email),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.SendKeys(`#userPswd`, config.Pass),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.Evaluate(`const e = document.querySelector("#privacyLabel"); if(e) { e.click() }`, nil),
		chromedp.Sleep(100*time.Millisecond),
		chromedp.WaitVisible(`#login-btn`),
		// chromedp.Submit(`#userAcct, #userPswd, #login-btn`),
		chromedp.Click(`#login-btn`),
		chromedp.Sleep(8*time.Second),
		chromedp.CaptureScreenshot(&screenshot),
		chromedp.WaitVisible(`.plant-data > .item:nth-child(2) > .data`),
		chromedp.Text(`.plant-data > .item:nth-child(2) > .data`, &kwh),
		chromedp.Sleep(time.Hour),
	)
	if err != nil {
		return err
	}
	kwh = strings.TrimSuffix(strings.TrimSpace(kwh), " kW")
	fmt.Println(kwh)
	return nil
}
