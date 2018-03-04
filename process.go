package main

import (
	"context"
	"fmt"

	"github.com/sclevine/agouti"
)

const (
	startPage   = 8
	seleniumURL = "http://localhost:4444/wd/hub"
	waitPerPage = 600000 // 10 minutes
)

func process(ctx context.Context) error {
	p, err := page(ctx)
	if err != nil {
		return err
	}
	for i := startPage; i >= 1; i-- {
		if err := p.Navigate(getURL(i)); err != nil {
			return err
		}
		links, err := p.FindByXPath("//img[@alt='Click here to download']/..").
			Elements()
		if err != nil {
			return err
		}
		for j, link := range links {
			if err := link.Click(); err != nil {
				return err
			}
		}
	}
	return nil
}

func getURL(p int) string {
	return fmt.Sprintf("https://interfacelift.com/wallpaper/downloads/date/widescreen_16:9/5120x2880/index%d.html", p)
}

func page(ctx context.Context) (*agouti.Page, error) {
	driver := agouti.Selenium()
	if err := driver.Start(); err != nil {
		return err
	}
	driver.WebDriver.Open()
	/*
		c := agouti.NewCapabilities().
			Browser("chrome").
			Platform("linux").
			With("javascriptEnabled")
		p, err := agouti.NewPage(seleniumURL, agouti.Desired(c))
		if err != nil {
			return nil, fmt.Errorf("error in NewPage: %v", err)
		}
	*/
	if err := p.SetImplicitWait(waitPerPage); err != nil {
		return nil, fmt.Errorf("error in SetImplicitWait: %v", err)
	}
	return p, nil
}
