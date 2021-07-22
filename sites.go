package main

import "github.com/Schmenn/sherlock/sites"

var siteList = map[string]func(username string) bool{
	"https://github.com/":                        sites.GitHub,
	"https://www.reddit.com/user/":               sites.Reddit,
	"https://osu.ppy.sh/users/":                  sites.Osu,
	"https://gitlab.com/":                        sites.GitLab,
	"https://instagram.com/":                     sites.Instagram,
	"https://tiktok.com/@":                       sites.TikTok,
	"https://www.chess.com/member/":              sites.Chess,
	"http://ws2.kik.com/user/":                   sites.Kik,
	"https://steamcommunity.com/id/":             sites.Steam,
	"https://twitch.tv/":                         sites.Twitch,
	"https://archive.org/details/@":              sites.ArchiveOrg,
	"https://hub.docker.com/u/":                  sites.DockerHub,
	"https://ifttt.com/p/":                       sites.IFTTT,
	"https://{u}.newgrounds.com":                 sites.NewGrounds,
	"https://www.pinterest.com/":                 sites.Pinterest,
	"https://pypi.org/user/":                     sites.PyPi,
	"https://www.roblox.com/user.aspx?username=": sites.Roblox,
	"https://{u}.slack.com":                      sites.Slack,
	"https://soundcloud.com/":                    sites.SoundCloud,
	"https://sourceforge.net/u/":                 sites.SourceForge,
	"https://open.spotify.com/user/":             sites.Spotify,
	"https://t.me/":                              sites.Telegram,
	"https://www.buzzfeed.com/":                  sites.BuzzFeed,
	"https://www.facebook.com/":                  sites.FaceBook,
	"https://euw.op.gg/summoner/userName=":       sites.OPGG,
}
