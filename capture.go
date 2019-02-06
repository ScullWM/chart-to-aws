package main

import (
	"github.com/chromedp/chromedp"
	"io/ioutil"

	"context"
	"fmt"
	"log"
	"time"
)

func capture(r string, q string, s string) {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var buf []byte
	url := fmt.Sprintf("%s%s", screenConfig.DomainScope, r)

	err = c.Run(ctxt, screenshot(url, s, &buf))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(q, buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(1 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.ByID),
	}
}
