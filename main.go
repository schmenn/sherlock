package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
)

func main() {
	// Colors used for output
	green := color.New(color.FgHiGreen)
	red := color.New(color.FgHiRed)

	// Exits when no username was provided
	if len(os.Args) < 2 {
		red.Println("Please specify a username")
		os.Exit(1)
	}

	// The Username
	name := os.Args[1]

	if strings.ContainsAny(name, "\t\u0020\n") {
		red.Println("Username cannot contain whitespace")
		os.Exit(1)
	}

	// Seeds gofakeit's randomizer
	gofakeit.Seed(time.Now().UnixNano())

	// Start all tasks concurrently
	var wg sync.WaitGroup
	for site, sitefunc := range siteList {
		wg.Add(1)
		go func(site string, sitefunc func(username string) bool, name string) {
			if sitefunc(name) {
				fmt.Println("[" + green.Sprint("+") + "] " + green.Sprint(makeURL(site, name)))
				// color.HiGreen("[+] " + makeURL(site, name))
			} else {
				fmt.Println("[" + red.Sprint("-") + "] " + red.Sprint(makeDomain(site)))
				// color.HiRed("[-] " + makeDomain(site))
			}
			wg.Done()
		}(site, sitefunc, name)

	}

	// Wait for all tasks to finish
	wg.Wait()
}

// makeURL creates a clickable URL when
// the user was found on the site
func makeURL(url, username string) string {
	if strings.Contains(url, "{u}") {
		return strings.ReplaceAll(url, "{u}", username)
	}
	return url + username
}

// makeDomain creates a domain when the
// user was not found on the site
func makeDomain(u string) string {
	u = strings.ReplaceAll(u, "{u}", "placeholder")
	m, _ := url.Parse(u)
	return strings.ReplaceAll(m.Hostname(), "placeholder.", "")
}
