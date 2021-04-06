package main

import (
	"os"
	"strings"
	"sync"

	"github.com/Schmenn/sherlock/sites"
	"github.com/fatih/color"
)

var siteList = map[string]func(username string) bool{
	"https://github.com/":                       sites.GitHub,
	"https://www.reddit.com/user/":              sites.Reddit,
	"https://osu.ppy.sh/users/":                 sites.Osu,
	"https://gitlab.com/api/v4/users?username=": sites.GitLab,
	"https://instagram.com/":                    sites.Instagram,
	"https://tiktok.com/@":                      sites.TikTok,
	"https://www.chess.com/member/":             sites.Chess,
	"http://ws2.kik.com/user/":                  sites.Kik,
	"https://steamcommunity.com/id/":      sites.Steam,
	"https://twitch.tv/":                        sites.Twitch,
}

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
				color.HiGreen("[+] "+ site + name)
			} else {
				color.HiRed("[-] " + strings.Split(strings.Split(site, "//")[1], "/")[0])
			}
			wg.Done()
		}(site, sitefunc, name)

	}
	wg.Wait()
}
