package main

import "github.com/Schmenn/sherlock/sites"

var siteList = map[string]func(username string) bool{
	"https://github.com/":                       sites.GitHub,
	"https://www.reddit.com/user/":              sites.Reddit,
	"https://osu.ppy.sh/users/":                 sites.Osu,
	"https://gitlab.com/api/v4/users?username=": sites.GitLab,
	"https://instagram.com/":                    sites.Instagram,
	"https://tiktok.com/@":                      sites.TikTok,
	"https://www.chess.com/member/":             sites.Chess,
	"http://ws2.kik.com/user/":                  sites.Kik,
	"https://steamcommunity.com/id/":            sites.Steam,
	"https://twitch.tv/":                        sites.Twitch,
	"https://archive.org/details/@":             sites.ArchiveOrg,
	"https://hub.docker.com/u/":                 sites.DockerHub,
	"https://ifttt.com/p/":                      sites.IFTTT,
	"https://{u}.newgrounds.com":                sites.NewGrounds,
}
