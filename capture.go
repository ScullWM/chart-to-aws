package main

import (
	"io/ioutil"

	"github.com/chromedp/chromedp"

	"context"
	"fmt"
	"time"
)

func capture(ctxt context.Context, query string, sel string, out string) error {
	var err error

	// create chrome instance
	c, err := chromedp.New(ctxt)
	if err != nil {
		return err
	}

	// run task list
	var buf []byte
	url := fmt.Sprintf("%s%s", screenConfig.DomainScope, query)

	err = c.Run(ctxt, screenshot(url, sel, &buf))
	if err != nil {
		return err
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		return err
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(out, buf, 0644)
	if err != nil {
		return err
	}

	return nil
}

func screenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(1 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.ByID),
	}
}
