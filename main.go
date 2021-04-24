package main

import (
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		color.HiRed("Please specify a username")
		os.Exit(1)
	}
	name := os.Args[1]
	var wg sync.WaitGroup
	for site, sitefunc := range siteList {
		wg.Add(1)
		go func(site string, sitefunc func(username string) bool, name string) {
			if sitefunc(name) {
				color.HiGreen("[+] " + makeURL(site, name))
			} else {
				color.HiRed("[-] " + makeDomain(site))
			}
			wg.Done()
		}(site, sitefunc, name)

	}
	wg.Wait()
}

func makeURL(url, username string) string {
	if strings.Contains(url, "{u}") {
		return strings.ReplaceAll(url, "{u}", username)
	}
	return url + username
}

func makeDomain(url string) string {
	return strings.Trim(
		strings.ReplaceAll(
			strings.Split(
				strings.Split(
					url, "//")[1],
				"/")[0],
			"{u}",
			"",
		),
		".",
	)
}
