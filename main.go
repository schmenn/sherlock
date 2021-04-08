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
				color.HiGreen("[+] " + site + name)
			} else {
				color.HiRed("[-] " + strings.Split(strings.Split(site, "//")[1], "/")[0])
			}
			wg.Done()
		}(site, sitefunc, name)

	}
	wg.Wait()
}
